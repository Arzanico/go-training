package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/google/uuid"
)

//minimalistic worker pool, I would like to implement something clean like this for my solution. This is still in progress

// task definition
type claimRequest struct {
	id   int
	body []byte
}

// way to process the task
func (o *claimRequest) process() {

	partnerUrl := "/partner/claims"
	request := http.Request{
		Method: "POST",
		URL:    &url.URL{Path: partnerUrl},
		Header: nil,
		Body:   io.NopCloser(bytes.NewBuffer(o.body)),
	}

	request.Header.Add("x-idempotency-key", uuid.NewString())
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	response, err := httpClient.Do(&request)
	if err != nil {
		fmt.Println(err)
	}

	responsePayload, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("response status code: ", response.StatusCode)
	fmt.Println("response body: ", string(responsePayload))
	defer func() {
		err = response.Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	fmt.Printf("claim request with id %d succesfully processed\n", o.id)

}

// worker pool definition
type workerPool struct {
	maxCocurrentWorkers int
	tasks               []claimRequest
	taskCh              chan claimRequest
	wg                  sync.WaitGroup
}

// function to execute the worker pool
func (o *workerPool) worker() {
	for task := range o.taskCh {
		task.process()
		o.wg.Done()
	}
}

func (o *workerPool) start() {
	//initialize the taks/jobs channel
	o.taskCh = make(chan claimRequest, len(o.tasks))

	//launch the go workers
	for i := 0; i < o.maxCocurrentWorkers; i++ {
		go o.worker()
	}

	//add all tasks to the waiting group
	o.wg.Add(len(o.tasks))

	//send tasks to the channel
	for _, t := range o.tasks {
		o.taskCh <- t
	}
	//close the channnel when all task are sent
	close(o.taskCh)

	//wait for all workers to finish with all tasks
	o.wg.Wait()
}
