package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	req, err := http.Get("http://www.google.com")
	if err != nil {
		panic(err)
	}
	fmt.Println("Status code", req.StatusCode)
	fmt.Println("Status code", req.Status)

	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))

	req.Body.Close()
}
