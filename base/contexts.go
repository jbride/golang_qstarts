package main

import (
	"fmt"
	"net/http"
	"time"
)

/*  Open two terminals:
 *    1) go run contexts.go
 *    2) curl localhost:20090/hello
 *       - wait 10 seconds for a response from server
 *       - <crtl-c> : observe behavior of ctx.Done event
 *
 */
func hello(rWriter http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(rWriter, "world\n")

	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(rWriter, err.Error(), internalError)
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":20090", nil)
}
