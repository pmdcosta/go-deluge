package deluge

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

// Client represents a client for managing Deluge RPC requests.
type Client struct {
	// deluge daemon auth.
	Url      string
	Password string

	// http client.
	client  *http.Client
	cookies []*http.Cookie

	// request counter.
	id int

	// client locker.
	lock sync.Mutex

	// service for for interacting with deluge.
	Service Service
}

// NewClient returns a new deluge client.
func NewClient(url string, password string) *Client {
	c := &Client{
		Url:      url,
		Password: password,
		client:   new(http.Client),
		cookies:  nil,
		id:       0,
	}
	c.Service.client = c

	return c
}

// Open starts the client and connects to the Deluge server.
func (c *Client) Open() error {
	response, err := c.sendRequest(AUTH, c.Password)
	if err != nil {
		return err
	}

	if response["result"] != true {
		return ErrAuthFailed
	}

	return nil
}

// Close stops the client.
func (c *Client) Close() error {
	return nil
}

// sendRequest makes an HTTP request to the Deluge server.
func (c *Client) sendRequest(method Method, params ...interface{}) (map[string]interface{}, error) {
	c.lock.Lock()
	defer c.lock.Unlock()

	// set request id.
	c.id++

	// build request data.
	data, err := json.Marshal(map[string]interface{}{
		"method": method,
		"id":     c.id,
		"params": params,
	})
	if err != nil {
		return nil, err
	}

	// build http request.
	req, err := http.NewRequest("POST", c.Url, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	if c.cookies != nil {
		for _, cookie := range c.cookies {
			req.AddCookie(cookie)
		}
	}

	// make http request.
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// parse response.
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("invalid status code: %d", resp.StatusCode)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// save cookies.
	c.cookies = resp.Cookies()

	// build request result.
	result := make(map[string]interface{})
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	if result["error"] != nil {
		return nil, fmt.Errorf("json error: %v", result["error"])
	}

	return result, nil
}
