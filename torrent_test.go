package deluge_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestClient_Auth tests connecting and authenticating a client
func TestClient_Auth(t *testing.T) {
	c := MustOpenClient()
	defer c.Close()
	assert.Nil(t, c.Auth())
}

// TestService_AddTorrentMagnet tests adding a torrent magnet.
func TestService_AddTorrentMagnet(t *testing.T) {
	c := MustOpenClient()
	defer c.Close()
	assert.Nil(t, c.Auth())

	_, err := c.AddTorrentMagnet("", nil)
	assert.Nil(t, err)
}

// TestService_AddTorrentUrl tests adding a torrent url.
func TestService_AddTorrentUrl(t *testing.T) {
	c := MustOpenClient()
	defer c.Close()
	assert.Nil(t, c.Auth())

	_, err := c.AddTorrentUrl("http://releases.ubuntu.com/18.10/ubuntu-18.10-desktop-amd64.iso.torrent", nil)
	assert.Nil(t, err)
}
