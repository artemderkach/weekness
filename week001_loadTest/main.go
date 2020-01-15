package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/pkg/errors"
)

func main() {
	c := make(chan time.Duration)

	reqNum := 10000
	go func(reqNum int) {
		var wg sync.WaitGroup

		for i := 0; i < reqNum; i += 1 {
			wg.Add(1)

			go func() {
				defer wg.Done()

				duration, err := request()
				if err != nil {
					log.Println(errors.Wrap(err, "error while request"))
					return
				}

				c <- duration
			}()
		}

		wg.Wait()
		close(c)
	}(reqNum)

	var count int64
	var sum time.Duration
	for duration := range c {
		fmt.Println(duration)
		count += 1
		sum += duration
	}

	fmt.Println("number of requests: ", count)
	fmt.Println("total request time: ", sum)
	fmt.Println("time per request:   ", time.Duration(sum.Nanoseconds()/count))
}

func request() (time.Duration, error) {
	req, err := http.NewRequest("GET", "http://localhost:8080", nil)
	if err != nil {
		return 0, errors.Wrap(err, "error creating request")
	}

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
