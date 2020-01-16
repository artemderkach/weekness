package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/jessevdk/go-flags"
	"github.com/pkg/errors"
)

type options struct {
	Protocol   string `short:"p" long:"protocol" default:"http" description:"http or https protocol"`
	Method     string `short:"m" long:"method" default:"GET" description:"http method"`
	Duration   int    `short:"d" long:"duration" default:"3" description:"requests duration (in seconds)"`
	RPS        int    `short:"r" long:"rps" default:"100" description:"requests per second"`
	Positional struct {
		URL string
	} `positional-args:"true" required:"true"`
}

type msg struct {
	duration time.Duration
	error    error
}

func main() {
	opts := &options{}
	_, err := flags.Parse(opts)
	if err != nil {
		// "--help" error
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		}

		os.Exit(1)
		// log.Fatal(errors.Wrap(err, "error parsing flags"))
	}
	// some validation
	if opts.Duration <= 0 {
		log.Fatal(errors.New("duration should be > 0"))
	}

	// create request from flags
	req, err := http.NewRequest(opts.Method, opts.Protocol+"://"+opts.Positional.URL, nil)
	if err != nil {
		log.Fatal(errors.Wrap(err, "error creating request"))
	}

	c := make(chan *msg)
	go runner(c, req, opts.RPS, opts.Duration)

	// calculate results
	var count int64
	var errCount int64
	var sum time.Duration
	for m := range c {
		count += 1
		if m.error != nil {
			errCount += 1
			continue
		}
		sum += m.duration
	}

	var duration time.Duration
	if count != 0 {
		duration = time.Duration(sum.Nanoseconds() / count)
	}
	fmt.Println("number of requests: ", count)
	fmt.Println("number of errors:   ", errCount)
	fmt.Println("time per request:   ", duration)
}

// runner makes multiple requests
func runner(c chan *msg, req *http.Request, rps int, duration int) {
	var wg sync.WaitGroup

	// "request per second" to "time used for one request"
	// nanoseconds per request (request interval in nanoseconds)
	npr := time.Duration(1000000000 / rps)
	timeShouldPass := npr

	start := time.Now()
	for i := 0; i < rps*duration; i += 1 {
		wg.Add(1)

		go func() {
			defer wg.Done()

			duration, err := request(req)
			m := &msg{
				duration: duration,
				error:    errors.Wrap(err, "error while request"),
			}

			c <- m
		}()

		timeShouldPass += npr
		timePassed := time.Since(start)
		timeToSleep := timeShouldPass - timePassed
		// case when request took longer than needed for required requests per second
		if timeToSleep <= 0 {
			continue
		}

		time.Sleep(timeToSleep)
	}

	wg.Wait()
	close(c)
}

// request
func request(req *http.Request) (time.Duration, error) {
	client := http.Client{
		Timeout: 3 * time.Second,
	}

	// time start
	startTime := time.Now()

	res, err := client.Do(req)
	if err != nil {
		return 0, errors.Wrap(err, "error sending request")
	}

	// time stop
	endTime := time.Now()

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, errors.Wrap(err, "error reading response")
	}

	return endTime.Sub(startTime), nil
}
