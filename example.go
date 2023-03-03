// Filename: go-httpclient/example.go
package go_httpclient
package main

import "github.com/JWEmmanuel89/go-httpclient/gohttp"
import (
	"fmt"
	"io"
)

func exampleUsage() {
	client := gohttp.New()
	client.Get()
	"github.com/JWEmmanuel89/go-httpclient/gohttp"
}

func main() {
	url := "https://api.github.com"
	client := gohttp.New()
	res, err := client.Get(url, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.StatusCode)
	defer res.Body.Close()
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}