package deluge_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/pmdcosta/go-deluge"
)

// Client is a test wrapper.
type Client struct {
	*deluge.Client
	Env Env
}

// Env holds env variables.
type Env struct {
	Url      string `json:"url"`
	Password string `json:"password"`
	Magnet   string `json:"magnet"`
	Torrent  string `json:"torrent"`
}

// NewClient returns a new instance of the wrapper Client.
func NewClient() *Client {
	// load env.
	var e Env
	raw, err := ioutil.ReadFile(".env")
	if err != nil {
		panic(err.Error())
	}
	json.Unmarshal(raw, &e)

	// wrap client.
	c := &Client{
		Client: deluge.NewClient(e.Url, e.Password),
		Env:    e,
	}
	return c
}

// MustOpenClient returns an new, open instance of Client.
func MustOpenClient() *Client {
	c := NewClient()
	if err := c.Client.Open(); err != nil {
		panic(err)
	}
	return c
}

// Close closes the client and removes the underlying database.
func (c *Client) Close() error {
	return c.Client.Close()
}

// TestClient_Auth tests connecting and authenticating a client.
func TestClient_Auth(t *testing.T) {
	c := MustOpenClient()
	defer c.Close()
}

// TestClient_Auth_Invalid_URL tests connecting and authenticating a client with an invalid URL.
func TestClient_Auth_Invalid_URL(t *testing.T) {
	c := deluge.NewClient("test", "test")
	if err := c.Open(); err == nil {
		t.Fatalf("connection should have failed")
	}
}

// TestClient_Auth_Invalid_PWD tests connecting and authenticating a client with an invalid password.
func TestClient_Auth_Invalid_PWD(t *testing.T) {
	c := deluge.NewClient("pmdcosta.powertrip.pt", "test")
	if err := c.Open(); err == nil {
		t.Fatalf("connection should have failed")
	}
}
