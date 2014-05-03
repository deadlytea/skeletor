package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Request struct {
	header  string
	message string
	code    int
}

func NewRequest(h string, m string, c int) *Request {
	req := new(Request)

	req.header = h
	req.message = m
	req.code = c

	return req
}

// Parses a request
// Adds to channel if request is valid, exits otherwise
func ParseRequest(req string, cProc chan *Request) {

	time.Sleep(250)

	request := strings.Split(req, '\n')

	// Check here if request is valid, return if not
	if len(request) != 3 {
		fmt.Printf("Bad request, discarding...\n")
		return
	}

	h := request[0]
	m := request[1]
	c, err := strconv.Atoi(request[2])

	if err != nil {
		fmt.Printf("Invalid request code, discarding...\n")
		return
	}

	// Add new request object to channel
	fmt.Printf("Request is valid, adding to channel to be processed\n")

	cProc <- NewRequest(h, m, c)

}

// Server loop to wait for requests
func GetRequests(cProc chan *Request) {

	fmt.Printf("Starting main loop...\n")

	for {

		// Wait for request
		time.Sleep(500)

		// TODO: Get the request

		fmt.Printf("Request received, sending to be parsed.\n")
		go ParseRequest(cProc)
	}

}

func Handler(cProc chan *Request) {

	fmt.Printf("Starting handler loop...\n")

	// TODO: Add logic to slow handling of requests for rate limited API's
	var req string
	for {
		req = <-cProc
		fmt.Printf("Parsed request taken from channel\n")
		go HandleRequest(req)
	}
}

// Request handler
// Using a plain string for now
func HandleRequest(req *Request) {
	fmt.Printf("Handling request...\n")
	time.Sleep(1000)
	fmt.Printf("Request handled\n")
}

func main() {
	cProcess := make(chan *Request, 50)

	fmt.Printf("Start me up!\n")

	// Start the server loop
	go GetRequests(cProcess)

	// Handle requests
	Handler(cProcess)

}
