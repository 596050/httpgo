package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/596050/httpgo/httpgo"
)

func getClientGithub() httpgo.Client {

	commonHeaders := make(http.Header)
	client := httpgo.NewBuilder().
		SetConnectionTimeout(2 * time.Second).
		SetResponseTimeout(5 * time.Second).
		SetHeaders(commonHeaders).
		Build()

	return client
}

var (
	httpClientGithub = getClientGithub()
)

// get
func getGithubUrls() {
	resp, err := httpClientGithub.Get("https://api.github.com", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(resp.Body()))
}

// post
type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func createUser(user User) {
	headers := make(http.Header)
	headers.Set("Authorization", "Bearer ABC-123")
	resp, err := httpClientGithub.Post("https://api.github.com", headers, user)
	if err != nil {
		panic(err)
	}
	var data User
	if err := resp.UnmarshalJSON(&data); err != nil {
		panic(err)
	}
	fmt.Println(data)
}

func main() {
	getGithubUrls()
}
