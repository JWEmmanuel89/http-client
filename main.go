// Filename: main.go
package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// to where are we making the request
	url := "https://api.github.com"
	// create a request
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}
	// Let's create a client
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)
	fmt.Println(string(bytes))

}
