package main

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/http/httptest"
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
	//server es un mock para testear el caso de uso del pool, en la proxima iteracion esto deberia estar afuera del process
	// para que todos los workeres le peguen al mismo servidor
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond)
			if r.URL.Path != partnerUrl {
				w.WriteHeader(http.StatusNotFound)
				_, _ = w.Write([]byte(`{"payload":"not found"}`))
				return
			}

			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"payload":"successfully processed"}`))

		}))
	defer server.Close()

	//mocked server url es la url aleatoria en donde se levanta el servidor de test + el path que yo quiero
	mockedServerUrl := server.URL + partnerUrl
	request, err := http.NewRequest(
		http.MethodPost,
		mockedServerUrl,
		bytes.NewBuffer(o.body),
	)

	if err != nil {
		panic(err)
	}

	request.Header.Set("x-idempotency-key", uuid.NewString())
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	response, err := httpClient.Do(request)
	defer func() {
		_ = response.Body.Close()
	}()
	if err != nil {
		panic(err)
	}

	responsePayload, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("response status code: ", response.StatusCode)
	fmt.Println("response body: ", string(responsePayload))
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
