// Web service for Fibonacci numbers - 3/24/16 Kam Pang
// - Accepts a number n as input and returns the first n Fibonacci numbers, starting from 0.
//   e.g. given n = 5, it will output [0, 1, 1, 2, 3], in JSON format.
// If given 0, 1, or 2, it will output the first two numbers [0, 1].
// If given a negative number or invalid entry which is not convertible to an integer:
// - It will respond with an error message, {"Error":"Input must be a positive integer"}
// To build, from a Linux host, such as Ubuntu 14.04:
// - verify go is installed, "go".
// - Set GOPATH to the working directory, "export GOPATH=`pwd`.
// - Enter "go build fibon_ws.go".
// To deploy/run:
// - Enter "./fibon_ws".
// Open a Web Browser, enter URL "locahost:8080/x", where x is the Fibonacci number count.
// - It will display the Fibonacci numbers.
// From a client terminal, enter "curl -X GET http://localhost:8080/x".
// - It will display the Fibonacci numbers.
// - To diplay the details and status code, use the "-vvv" option.
// To manually test:
// - Enter positive integers for x, such as 3, 12.
// - Enter negative numbers for x, such as -5, -10.
// - Enter invalid values, which cannot be converted to integers, such as xyz, 5a.
// - Verfiy the output are as expected.
// To perform functional test:
// - From a client terminal, run the script fibon_ws_test.sh in the test directory.
// To perform unit test:
// - fibon_ws_test.go needs to be in the same directory as fibon_ws.go.
// - Enter "go test".
// To exit the program: enter "Ctrl c".

package main

import (
  "fmt"
  "net/http"
  "strconv"
  "encoding/json"
)

type Error_message struct {
  Error string
}

func Calculate_sequence(input_number int) []int {

  buf := [1024]int {0,1} // initialize an 1024-length array of integers populated with [0, 1, 0(uninitialized), 0(uninitialized), ...]

  i := 2 // start at 2 since we have already put 2 elements in the array above
  // Fn = Fn-1 + Fn-2
  for i < input_number {
    buf[i] = buf[i - 1] + buf[i - 2] // buffer for current iteration equals buffer from two iterations ago + buffer from 1 iteraton ago
    i++
  }

  return buf[0:i]  // slice of the buffer with elements zero through our max Fibonacci iteration

}

func fibon_handler(writer http.ResponseWriter, request *http.Request) {

  input_number, err := strconv.Atoi(request.URL.Path[1:]) // input_number from the second slice of the URL path onward
  writer.Header().Set("Content-Type", "application/json; charset=utf-8") // set response content type

  // if the request is for less than 0 or something that cannot be converted to an integer, send client an error message
  if err != nil || input_number < 0 {
    the_error := Error_message{"Input must be a positive integer"}
    the_error_json_bytes, _ := json.Marshal(the_error)
    the_error_json_string := string(the_error_json_bytes)
    writer.WriteHeader(http.StatusBadRequest) // status = 400 Bad Request
    fmt.Fprint(writer, the_error_json_string, "\n") 
    return
  }

  response_json_bytes, _ := json.Marshal(Calculate_sequence(input_number))
  response_json_string := string(response_json_bytes)
  // default status = 200 OK
  fmt.Fprint(writer, response_json_string, "\n")

  return

}

func main() {
  http.HandleFunc("/", fibon_handler) // use fibon_handler to handle GET requests 
  http.ListenAndServe(":8080", nil)
}
