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

type Options struct {
	Protocol   string `short:"p" long:"protocol" default:"http" description:"http or https protocol"`
	Method     string `short:"m" long:"method" default:"GET" description:"http method"`
	ReqNum     int    `short:"n" long:"reqNum" default:"100" description:"number of requests"`
	Positional struct {
		URL string
	} `positional-args:"true" required:"true"`
}

func main() {
	opts := &Options{}
	_, err := flags.Parse(opts)
	if err != nil {
		// "--help" error
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		}

		os.Exit(1)
		// log.Fatal(errors.Wrap(err, "error parsing flags"))
	}

	// create request from flags
	reqNum := opts.ReqNum
	req, err := http.NewRequest(opts.Method, opts.Protocol+"://"+opts.Positional.URL, nil)
	if err != nil {
		log.Fatal(errors.Wrap(err, "error creating request"))
	}

	c := make(chan time.Duration)
	go runner(c, req, reqNum)

	// calculate results
	var count int64
	var sum time.Duration
	for duration := range c {
		fmt.Println(duration)
		count += 1
		sum += duration
	}

	fmt.Println("number of requests: ", count)
	fmt.Println("time per request:   ", time.Duration(sum.Nanoseconds()/count))
}

// runner makes multiple requests
func runner(c chan time.Duration, req *http.Request, reqNum int) {
	var wg sync.WaitGroup

	for i := 0; i < reqNum; i += 1 {
		wg.Add(1)

		go func() {
			defer wg.Done()

			duration, err := request(req)
			if err != nil {
				log.Println(errors.Wrap(err, "error while request"))
				return
			}

			c <- duration
		}()
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
