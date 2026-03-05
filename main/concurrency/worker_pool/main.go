package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type internalRequest struct {
	id int
}

type externalResponse struct {
	idempotencyKey int
	statusCode     int
	body           []byte
}

func (o *externalResponse) reply() {
	// random time for sending the externalResponse back
	//time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	n := rand.Intn(10)
	var responseCode int
	if n > 5 {
		responseCode = 500
	} else {
		responseCode = 400
	}

	o.statusCode = responseCode
	o.body = []byte("hi")
}

func endpointRequests() []internalRequest {
	// this simulates the internalRequest from an endpoint
	listOfRequests := make([]internalRequest, 0)
	for i := 0; i < 40; i++ {
		r := internalRequest{
			id: i,
		}
		listOfRequests = append(listOfRequests, r)
	}
	return listOfRequests
}

func validateRequest(r internalRequest) bool {
	// to generate some noice we will have an invalid internalRequest for each prime reques id
	for i := 2; i < 10; i++ {
		if r.id%i == 0 {
			return true
		}
	}
	return false
}

func enrichData(data string) string {
	return data
}

type storage struct {
	requests map[int]*internalRequest
	response map[int]*externalResponse
}

func (s *storage) storeRequest(request internalRequest) {
	s.requests[request.id] = &request
}

func (s *storage) storeResponse(response externalResponse) {
	s.response[response.idempotencyKey] = &response
}

func main() {

	//internal request represents a request from an internal user, this request will then be validate and passed to a partner
	requests := endpointRequests()
	claims := make([]internalRequest, 0)
	for _, request := range requests {
		if !validateRequest(request) {
			fmt.Printf("Invalid request: %v\n", request.id)
			continue
		}
		claims = append(claims, request)
	}

	storageOb := &storage{
		requests: make(map[int]*internalRequest),
		response: make(map[int]*externalResponse),
	}

	handleClaims(storageOb, claims)

}

type claim struct {
	index int
	internalRequest
}

func (c *claim) handleClaim(respCh chan externalResponse) {
	exResponse := externalResponse{
		idempotencyKey: c.id,
	}
	exResponse.reply()
	respCh <- exResponse
}

func worker(cCh chan *claim, resp chan externalResponse, wg *sync.WaitGroup) {
	for task := range cCh {
		task.handleClaim(resp)
		wg.Done()
	}

}

func handleClaims(_ *storage, r []internalRequest) {
	jobChan := make(chan *claim, len(r))
	responseChannel := make(chan externalResponse, len(r))
	partnerResponse := make([]externalResponse, 0, len(r))

	wg := sync.WaitGroup{}
	wg.Add(len(r))

	for i := 0; i < 20; i++ {
		go worker(jobChan, responseChannel, &wg)
	}

	for i, iR := range r {
		c := &claim{
			index:           i,
			internalRequest: iR,
		}
		jobChan <- c

	}
	close(jobChan)

	wg.Wait()
	close(responseChannel)

	for result := range responseChannel {
		partnerResponse = append(partnerResponse, result)

	}

	for _, response := range partnerResponse {
		fmt.Printf("Partner response: %d\n", response.idempotencyKey)
		fmt.Printf("Partner response: %d\n", response.statusCode)
	}
}
