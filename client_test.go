package deluge_test

import (
	"github.com/pmdcosta/go-deluge"
	"github.com/sirupsen/logrus"
)

// Client is a test wrapper.
type Client struct {
	*deluge.Client
}

// NewClient returns a new instance of the wrapper Client.
func NewClient() *Client {
	// wrap client.
	c := &Client{
		Client: deluge.NewClient("http://127.0.0.1:8112/json", "deluge", logrus.New()),
	}
	return c
}

// MustOpenClient returns an new, open instance of Client.
func MustOpenClient() *Client {
	c := NewClient()
	c.Client.Open()
	return c
}

// Close closes the client and removes the underlying database.
func (c *Client) Close() {
	c.Client.Close()
}
