package typeform

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	baseURL = "https://api.typeform.com"
)

type Client struct {
	client *http.Client
	key    string
}

// Forms retrieves all forms
func (c *Client) Forms() (*FormsResponse, error) {
	url := fmt.Sprintf("%s/forms", baseURL)
	var forms FormsResponse
	err := c.executeRequest("GET", url, nil, &forms)
	return &forms, err
}

func (c *Client) FormByID(formID string) (*Form, error) {
	url := fmt.Sprintf("%s/forms/%s", baseURL, formID)
	var form Form
	err := c.executeRequest("GET", url, nil, &form)
	if err != nil {
		return nil, err
	}
	return &form, nil
}

func (c *Client) Responses(formID string, q *Query) (*ResponsesResponse, error) {
	var params string
	if q != nil {
		params = q.Encode()
	}
	url := fmt.Sprintf("%s/forms/%s/responses?%s", baseURL, formID, params)
	var resp ResponsesResponse
	err := c.executeRequest("GET", url, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) executeRequest(method, url string, r io.Reader, data interface{}) error {
	log.Println("Executing: ", url)
	req, err := c.newRequest("GET", url, nil)
	if err != nil {
		return err
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = readResponse(resp.Body, data)
	return err
}

func (c *Client) newRequest(method, url string, r io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, r)
	req.Header.Set("Authorization", "Bearer "+c.key)
	return req, err
}

func NewClient(key string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	c := &Client{
		key:    key,
		client: httpClient,
	}
	return c
}

func readResponse(r io.Reader, data interface{}) error {
	return json.NewDecoder(r).Decode(data)
}
