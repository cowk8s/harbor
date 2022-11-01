package http

import (
	"net/http"

	"github.com/cowk8s/harbor/src/common/http/modifier"
)

// Client is a util for common HTTP operations, such Get, Head, Post, Put and Delete.
// Use Do instead if those methods can not meet your requiremnet
type Client struct {
	modifiers []modifier.Modifier
	client    *http.Client
}

func (c *Client) Getclient() *http.Client {
	return c.client
}

func NewClient(c *http.Client, modifiers ...modifier.Modifier) *Client {
	client := &Client{
		client: c,
	}
	if client.client == nil {
		client.client = &http.Client{
			Transport: ,
		}
	}
}