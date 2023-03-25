package consumer

import (
	"anchor_exercise/model"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Consumer struct {
	BaseURL    *url.URL
	httpClient *http.Client
}

func (c *Consumer) GetNode(id int) (*model.Node, error) {
	req, err := c.newRequest("GET", fmt.Sprintf("/node/%d", id), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	var message model.Node
	_, err = c.do(req, &message)

	if err != nil {
		return nil, ErrUnavailable
	}

	return &message, err

}

func (c *Consumer) newRequest(method, path string, body interface{}) (*http.Request, error) {
	rel := &url.URL{Path: path}
	u := c.BaseURL.ResolveReference(rel)
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "Admin Service")

	return req, nil
}

func (c *Consumer) do(req *http.Request, v interface{}) (*http.Response, error) {
	if c.httpClient == nil {
		c.httpClient = http.DefaultClient
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}

var (
	ErrUnavailable = errors.New("api unavailable")
)
