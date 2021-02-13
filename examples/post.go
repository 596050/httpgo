package examples

import (
	"errors"
	"fmt"
	"net/http"
)

type GithubError struct {
	StatusCode       int    `json:"-"`
	Message          string `json:"message"`
	DocumentationURL string `json:"documentation_url"`
}

type Repository struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Private     bool   `json:"private"`
}

func CreateRepository(request Repository) (*Repository, error) {
	response, err := httpClient.Post("https://api.github.com/user/repos", request)

	fmt.Println("response, err", response, err)

	if err != nil {
		return nil, err
	}

	if response.StatusCode() != http.StatusCreated {
		var githubError GithubError
		if err := response.UnmarshalJSON(&githubError); err != nil {
			return nil, errors.New("error processing github response when creating a new repo")
		}
		return nil, errors.New(githubError.Message)
	}
	var result Repository
	if err := response.UnmarshalJSON(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
