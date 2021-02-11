package examples

import "fmt"

type Endpoints struct {
	CurrentUserURL    string `json:"current_user_url"`
	AuthorizationsURL string `json:"authorizations_url"`
	RepositoryURL     string `json:"repository_url"`
}

func GetEndpoints() (*Endpoints, error) {
	response, err := httpClient.Get("https://api.github.com", nil)
	if err != nil {
		return nil, err
	}
	fmt.Println(fmt.Sprintf("Status Code: %d", response.StatusCode()))
	fmt.Println(fmt.Sprintf("Status: %s", response.Status()))
	fmt.Println(fmt.Sprintf("Body: %s\n", string(response.Body())))
	var endpoints Endpoints
	if err := response.UnmarshalJSON(&endpoints); err != nil {
		return nil, err
	}
	fmt.Println(fmt.Sprintf("Repository URL: %s", endpoints.RepositoryURL))
	return &endpoints, nil
}
