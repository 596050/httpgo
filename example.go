package main

import (
	"fmt"
	"io/ioutil"

	"github.com/596050/httpgo/httpgo"
)

// func rootHandler(w http.ResponseWriter, r *http.Request) {
// 	time.Sleep(10 * time.Second)
// 	fmt.Fprint(w, "hello world!")
// }

// func serveRootHandler() {
// 	http.HandleFunc("/", rootHandler)
// 	http.ListenAndServe(":8080", nil)
// }

func main() {
	client := httpgo.New()
	resp, err := client.Get("https://api.github.com", nil)
	if err != nil {
		panic(err)
	}
	bytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bytes))
}
