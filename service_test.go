package deluge_test

import (
	"testing"
)

// TestService_AddTorrentMagnet tests adding a torrent magnet.
func TestService_AddTorrentMagnet(t *testing.T) {
	c := MustOpenClient()
	defer c.Close()
	_, err := c.Service().AddTorrentMagnet(c.Env.Magnet, nil)

	if err != nil {
		t.Fatalf("failed to add magnet: %v", err)
	}
}

// TestService_AddTorrentUrl tests adding a torrent url.
func TestService_AddTorrentUrl(t *testing.T) {
	c := MustOpenClient()
	defer c.Close()
	_, err := c.Service().AddTorrentUrl(c.Env.Torrent, nil)

	if err != nil {
		t.Fatalf("failed to add magnet: %v", err)
	}
}
