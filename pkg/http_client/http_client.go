package http_client

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

// HTTPClient
// manages to make http calls.
type HTTPClient struct {
	client *http.Client
}

// New
// creating a new http client.
func New() *HTTPClient {
	return &HTTPClient{
		client: &http.Client{},
	}
}

// Get
// making a get request.
func (h HTTPClient) Get(uri string, headers ...string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, fmt.Errorf("failed in creating requests: %w", err)
	}

	for _, pair := range headers {
		parts := strings.Split(pair, ":")

		req.Header.Add(parts[0], parts[1])
	}

	resp, err := h.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http request failed: %w", err)
	}

	return resp, nil
}

// Post
// making a post request.
func (h HTTPClient) Post(uri string, body io.Reader, headers ...string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPost, uri, body)
	if err != nil {
		return nil, fmt.Errorf("failed in creating requests: %w", err)
	}

	for _, pair := range headers {
		parts := strings.Split(pair, ":")

		req.Header.Add(parts[0], parts[1])
	}

	resp, err := h.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http request failed: %w", err)
	}

	return resp, nil
}

// Put
// making a put request.
func (h HTTPClient) Put(uri string, body io.Reader, headers ...string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPut, uri, body)
	if err != nil {
		return nil, fmt.Errorf("failed in creating requests: %w", err)
	}

	for _, pair := range headers {
		parts := strings.Split(pair, ":")

		req.Header.Add(parts[0], parts[1])
	}

	resp, err := h.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http request failed: %w", err)
	}

	return resp, nil
}

// Delete
// making a delete request.
func (h HTTPClient) Delete(uri string, headers ...string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodDelete, uri, nil)
	if err != nil {
		return nil, fmt.Errorf("failed in creating requests: %w", err)
	}

	for _, pair := range headers {
		parts := strings.Split(pair, ":")

		req.Header.Add(parts[0], parts[1])
	}

	resp, err := h.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http request failed: %w", err)
	}

	return resp, nil
}
