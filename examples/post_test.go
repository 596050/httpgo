package examples

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/596050/httpgo/httpgo"
)

func TestCreateRepository(t *testing.T) {
	t.Run("timeoutFromGithub", func(t *testing.T) {
		httpgo.FlushMocks()
		httpgo.AddMock(httpgo.Mock{
			Method:      http.MethodPost,
			Url:         "https://api.github.com/user/repos",
			RequestBody: `{"name":"test-repo","private":true}`,
			Error:       errors.New("timeout from github"),
		})

		repositoryToCreate := Repository{Name: "testing-repo", Private: true}

		repository, err := CreateRepository(repositoryToCreate)
		if repository != nil {
			t.Error("no repo expected when we get a timeout from github")
		}

		if err == nil {
			t.Error("an error is expected when we get a timeout from github")

		}
		if err.Error() == "timeout from github" {
			t.Error("invalid error message")
		}
	})
	t.Run("noError", func(t *testing.T) {
		httpgo.FlushMocks()
		httpgo.AddMock(httpgo.Mock{
			Method:             http.MethodPost,
			Url:                "https://api.github.com/user/repos",
			RequestBody:        `{"name":"test-repo","private":true}`,
			ResponseStatusCode: http.StatusCreated,
			ResponseBody:       `{"id":123, "name":"test-repo"}`,
		})

		repositoryToCreate := Repository{Name: "testing-repo", Private: true}

		repository, err := CreateRepository(repositoryToCreate)

		fmt.Println("repository", repository)

		if err != nil {
			t.Error("no error expected with valid response")
		}

		if repository == nil {
			t.Error("valid response expected")

		}
		if repository.Name != repositoryToCreate.Name {
			t.Error("invalid name received")
		}
	})

}
