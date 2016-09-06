package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	body := body_for_url("http://www.google.com/")
	fmt.Println(body)
}

func body_for_url(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("No good")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body)
}
