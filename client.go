package deluge

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/sirupsen/logrus"
)

// Client handles the connection and interface for managing Deluge RPC requests
type Client struct {
	logger *logrus.Entry

	// deluge daemon credentials
	Url      string
	Password string

	// underlying http client
	client  *http.Client
	cookies []*http.Cookie
	lock    sync.Mutex

	// deluge request counter
	id int
}

// NewClient instantiates a new client
func NewClient(url string, password string, logger *logrus.Logger) *Client {
	return &Client{
		Url:      url,
		Password: password,
		logger:   logger.WithFields(logrus.Fields{"client": "deluge-go"}),
	}
}

// Open creates the connection to the Deluge daemon
func (c *Client) Open() {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.client = new(http.Client)
	c.cookies = nil
	c.id = 0
}

// Close stops the connection to the Deluge daemon
func (c *Client) Close() {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.client = nil
	c.cookies = nil
	c.id = 0
}

// sendRequest makes an HTTP request to the Deluge daemon
func (c *Client) sendRequest(method Method, params ...interface{}) (map[string]interface{}, error) {
	c.lock.Lock()
	defer c.lock.Unlock()

	// increment request id
	c.id++

	// build http request body
	data, err := json.Marshal(map[string]interface{}{
		"method": method,
		"id":     c.id,
		"params": params,
	})
	if err != nil {
		return nil, err
	}

	// build http request
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

	// make http request
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// parse response
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("invalid status code: %d", resp.StatusCode)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// save cookies
	c.cookies = resp.Cookies()

	// build request result
	result := make(map[string]interface{})
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	// check if successful
	if result["error"] != nil {
		return nil, fmt.Errorf("json error: %v", result["error"])
	}

	return result, nil
}
