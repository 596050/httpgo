package examples

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
	var endpoints Endpoints
	if err := response.UnmarshalJSON(&endpoints); err != nil {
		return nil, err
	}
	return &endpoints, nil
}
