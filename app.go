package main

import "fmt"
import "time"

// Parses a request
// Adds to channel if request is valid, exits otherwise
func ParseRequest(cProc chan string) {

  time.Sleep(250)

  if true {
    defer fmt.Printf("Request is valid, adding to channel to be processed")
    cProc <- "String to process"
    return
  }

  fmt.Printf("Bad request, discarding...")

}

// Server loop to wait for requests
func GetRequests(cProc chan string) {

  fmt.Printf("Starting main loop...\n")

  for {
    // Wait for request
    time.Sleep(500)
    fmt.Printf("Request received, sending to be parsed.\n")
    go ParseRequest(cProc)
  }

}

func Handler(cProc chan string) {

  fmt.Printf("Starting handler loop...\n")

  // TODO: Add logic to slow handling of requests for rate limited API's
  var req string
  for {
    req = <- cProc
    fmt.Printf("Parsed request taken from channel\n")
    go HandleRequest(req)
  }
}

// Request handler
// Using a plain string for now
func HandleRequest(req string) {
  fmt.Printf("Handling request...\n")
  time.Sleep(1000)
  fmt.Printf("Request handled\n")
}

func main() {
  cProcess := make(chan string, 50)

  fmt.Printf("Start me up!\n")

  go GetRequests(cProcess)
  go Handler(cProcess)

  // wait?
  for {

  }

}
