package examples

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/596050/httpgo/httpgo"
)

func TestMain(m *testing.M) {
	fmt.Println("Testing package")
	// mock http requests
	httpgo.StartMockServer()
	os.Exit(m.Run())
}

func TestGetEndpoints(t *testing.T) {

	t.Run("TestErrorFetchingFromGithub", func(t *testing.T) {
		// Initialization
		httpgo.AddMock(httpgo.Mock{
			Method: http.MethodGet,
			Url:    "https://api.github.com",
			Error:  errors.New("timeout getting github endpoints"),
		})

		// Execution
		endpoints, err := GetEndpoints()

		// Validation
		if endpoints != nil {
			t.Error("no endpoints expected")
		}
		if err == nil {
			t.Error("error expected")
		}
		if err.Error() != "timeout getting github endpoints" {
			t.Error("invalid error message")
		}
	})
	t.Run("TestErrorUnmarshalResponseBody", func(t *testing.T) {
		// Initialization
		httpgo.AddMock(httpgo.Mock{
			Method: http.MethodGet,
			Url:    "https://api.github.com",
			ResponseBody: `{
            "current_user_url": 123,
            "authorizations_url": "https://api.github.com/authorizations",
            "repository_url": 123
        }`,
			ResponseStatusCode: http.StatusOK,
		})
		// Execution
		endpoints, err := GetEndpoints()
		// Validation
		if endpoints != nil {
			t.Error("no endpoints expected")
		}
		if err == nil {
			t.Error("error expected")
		}
		if !strings.Contains(err.Error(), "cannot unmarshal number into Go struct field") {
			t.Error(fmt.Sprintf("invalid error message, '%s'", err.Error()))
		}
	})
	t.Run("TestNoError", func(t *testing.T) {
		// Initialization
		httpgo.AddMock(httpgo.Mock{
			Method: http.MethodGet,
			Url:    "https://api.github.com",
			ResponseBody: `{
            "current_user_url": "https://api.github.com/user",
            "authorizations_url": "https://api.github.com/authorizations",
            "repository_url": "https://api.github.com/repos/{owner}/{repo}"
        }`,
			ResponseStatusCode: http.StatusOK,
		})
		// Execution
		endpoints, err := GetEndpoints()

		// Validation
		if err != nil {
			t.Error(fmt.Sprintf("no error expected, '%s'", err.Error()))
		}
		if endpoints == nil {
			t.Error("endpoints expected")
		}
		if endpoints.CurrentUserURL != "https://api.github.com/user" {
			t.Error("invalid current user URL")
		}
	})

}
