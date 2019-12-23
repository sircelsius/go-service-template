package foaas

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	_http "github.com/sircelsius/go-service-template/internal/http"
)

type Client struct {
	client *_http.Client
}

func NewClient() *Client {
	return &Client{
		client: _http.NewClient("foaas"),
	}
}

func (c *Client) Cool(ctx context.Context, from string) (string, error) {
	type httpResponse struct {
		Message string `json:"message"`
	}
	r, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://foaas.com/cool/%v", from), nil)
	if err != nil {
		return "", err
	}

	r.Header.Add("Accept", "application/json")

	reply, err := c.client.Do(ctx, r)

	if err != nil {
		return "", err
	}
	defer reply.Body.Close()
	var response httpResponse
	err = json.NewDecoder(reply.Body).Decode(&response)
	if err != nil {
		return "", err
	}

	return response.Message, nil
}
