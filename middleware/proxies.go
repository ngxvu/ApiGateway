package middleware

import (
	"net/http"
	"net/url"
)

func ProxyRequest(endpoint string, method string, queryParams map[string]string, apiKey string) (*http.Response, error) {
	client := &http.Client{}

	// Parse the endpoint URL
	parsedURL, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	// Add query parameters to the URL
	q := parsedURL.Query()
	for key, value := range queryParams {
		q.Add(key, value)
	}
	// Add API key as the last parameter
	q.Add("api_key", apiKey)
	parsedURL.RawQuery = q.Encode()

	req, err := http.NewRequest(method, parsedURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return client.Do(req)
}
