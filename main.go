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

func insert_user(username, password string) (user_id int, err error) {
	return 1, nil
}

// func get_user() {
// 	conn := util.Connect("connect test")
// 	defer conn.Close()
// 	var username string
// 	var password string
// 	var id int64 = 1
// 	var balance int64
// 	err := conn.QueryRow(`
//     select username,
//            password,
//            balance
//       from users
//      where id = $1`, id).Scan(&username, &password, &balance)
// 	if err != nil {
// 		if err == pgx.ErrNoRows {
// 			fmt.Fprintf(os.Stderr, "User with id %s not found.\n", id)
// 		} else {
// 			fmt.Fprintf(os.Stderr, "Unexpected error trying to find user: %v\n", err)
// 		}
// 		os.Exit(1)
// 	}

// 	fmt.Printf("Found user %s %s\n", username, password)
// 	fmt.Printf("Connection worked!\n")
// }

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
