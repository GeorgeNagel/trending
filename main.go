package main

import (
	"fmt"
	// "io/ioutil"
	"net/http"
)

type Meme struct {
	phrase string
	price  float32
}

var m1 = Meme{"harambe", 23.4}

func main() {
	http.HandleFunc("/", meme_handler)
	http.ListenAndServe(":8083", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func meme_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Phrase: %s. Price: $%v", m1.phrase, m1.price)
}

// // Get the content string at a url
// func body_for_url(url string) string {
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		fmt.Println("No good")
// 	}
// 	defer resp.Body.Close()
// 	body, err := ioutil.ReadAll(resp.Body)
// 	return string(body)
// }
