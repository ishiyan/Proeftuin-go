package assignment

import (
	"net/http"
)

// HttpClient allows to customize the underlying HTTP client.
type HttpClient interface {
	Do(request *http.Request) (*http.Response, error)
}

// AccountClientBuilder builds an account client.
type AccountClientBuilder interface {
	WithtHeaders(headers http.Header) AccountClientBuilder
	WithHttpClient(httpClient HttpClient) AccountClientBuilder

	Build() AccountClient
}

type builder struct {
	headers    http.Header
	baseUrl    string
	httpClient HttpClient
}

// NewAccountBuilder creates a new account client builder.
func NewAccountBuilder(baseUrl string) AccountClientBuilder {
	b := &builder{
		headers: http.Header{},
		baseUrl: baseUrl,
	}

	b.headers.Add("Content-Type", "application/vnd.api+json")

	return b
}

// Build builds a new account client.
func (b *builder) Build() AccountClient {
	c := &accountClient{
		httpClient: b.httpClient,
		baseUrl:    b.baseUrl,
		headers:    b.headers,
	}

	if c.httpClient == nil {
		c.httpClient = &http.Client{}
	}

	return c
}

// WithHeaders sets request headers.
func (b *builder) WithtHeaders(headers http.Header) AccountClientBuilder {
	b.headers = headers
	return b
}

// WithHttpClient sets an external http client.
func (b *builder) WithHttpClient(httpClient HttpClient) AccountClientBuilder {
	b.httpClient = httpClient
	return b
}
