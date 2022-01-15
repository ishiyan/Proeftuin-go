package assignment

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	errFmtUnexpectedStatusCode = "unexpected status code %d"
)

var (
	errAccountDoesNotExist    = errors.New("specified account ID does not exist")
	errAccounVersionIncorrect = errors.New("specified account version is incorrect")
)

// AccountClient implements functionality to create, fetch and delete accounts.
type AccountClient interface {
	Create(ctx context.Context, account *AccountData) (*AccountData, error)
	Fetch(ctx context.Context, accountID string) (*AccountData, error)
	Delete(ctx context.Context, accountID, version string) error
}

type accountClient struct {
	httpClient HttpClient
	baseUrl    string
	headers    http.Header
}

type response struct {
	statusCode int
	body       []byte
}

func (a *accountClient) Create(ctx context.Context, account *AccountData) (*AccountData, error) {
	url := a.baseUrl + "/v1/organisation/accounts"
	resp, err := a.do(ctx, http.MethodPost, url, account)
	if err != nil {
		return nil, err
	}

	switch resp.statusCode {
	case 201:
		break
	default:
		return nil, fmt.Errorf(errFmtUnexpectedStatusCode, resp.statusCode)
	}

	return unmarshalBody(resp.body)
}

func (a *accountClient) Fetch(ctx context.Context, accountID string) (*AccountData, error) {
	url := a.baseUrl + "/v1/organisation/accounts/" + accountID
	resp, err := a.do(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	switch resp.statusCode {
	case 200:
		break
	case 404:
		return nil, errAccountDoesNotExist
	default:
		return nil, fmt.Errorf(errFmtUnexpectedStatusCode, resp.statusCode)
	}

	return unmarshalBody(resp.body)
}

func (a *accountClient) Delete(ctx context.Context, accountID, version string) error {
	url := a.baseUrl + "/v1/organisation/accounts/" + accountID + "?version=" + version
	resp, err := a.do(ctx, http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	switch resp.statusCode {
	case 204:
		return nil
	case 404:
		return errAccountDoesNotExist
	case 409:
		return errAccounVersionIncorrect
	default:
		return fmt.Errorf(errFmtUnexpectedStatusCode, resp.statusCode)
	}
}

func (a *accountClient) do(ctx context.Context, method string, url string, body *AccountData) (*response, error) {
	reqBody, err := marshalBody(body)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header = a.headers

	resp, err := a.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	r := response{
		statusCode: resp.StatusCode,
		body:       respBody,
	}

	return &r, nil
}

func marshalBody(body *AccountData) (*bytes.Buffer, error) {
	if body == nil {
		return nil, nil
	}

	bs, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	return bytes.NewBuffer(bs), nil
}

func unmarshalBody(body []byte) (*AccountData, error) {
	var acc AccountData
	if err := json.Unmarshal(body, &acc); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return &acc, nil
}
