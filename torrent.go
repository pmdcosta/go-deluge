package deluge

import (
	"github.com/sirupsen/logrus"
)

// Auth establishes an authenticated connection with the Deluge daemon
func (c *Client) Auth() error {
	c.logger.Debug("connecting to Deluge daemon...")
	r, err := c.sendRequest(MethodAuth, c.Password)
	if err != nil {
		c.logger.WithFields(logrus.Fields{"url": c.Url, "err": err}).Error("failed to establish authentication request to the Deluge daemon.")
		return ErrRequestFailed
	}

	if r["result"] != true {
		c.logger.WithFields(logrus.Fields{"url": c.Url, "response": r}).Error("failed to authenticate with the Deluge daemon.")
		return ErrAuthFailed
	}

	return nil
}

// AddTorrentMagnet starts downloading a magnet link
// magnetUrl is the Magnet URL for the torrent
// options is a map with options to be set (check the Deluge Torrent documentation)
func (c *Client) AddTorrentMagnet(magnetUrl string, options map[string]interface{}) (string, error) {
	c.logger.WithFields(logrus.Fields{"url": magnetUrl, "options": options}).Debug("adding new torrent magnet...")
	if magnetUrl == "" {
		return "", ErrInvalidTorrent
	}

	response, err := c.sendRequest(MethodAddMagnet, magnetUrl, options)
	if err != nil {
		c.logger.WithFields(logrus.Fields{"url": c.Url, "err": err}).Error("failed to establish magnet creation request to the Deluge daemon.")
		return "", err
	}

	if response["result"] == nil {
		c.logger.WithFields(logrus.Fields{"url": c.Url, "err": err}).Error("failed to add torrent magnet.")
		return "", ErrAddTorrentFailed
	}

	return response["result"].(string), nil
}

// AddTorrentUrl starts downloading a torrent from a URL
// torrentUrl is the URL for the torrent
// options is a map with options to be set (check the Deluge Torrent documentation)
func (c *Client) AddTorrentUrl(torrentUrl string, options map[string]interface{}) (string, error) {
	c.logger.WithFields(logrus.Fields{"url": torrentUrl, "options": options}).Debug("adding new torrent url...")
	if torrentUrl == "" {
		return "", ErrInvalidTorrent
	}

	response, err := c.sendRequest(MethodAddTorrentURL, torrentUrl, options)
	if err != nil {
		c.logger.WithFields(logrus.Fields{"url": c.Url, "err": err}).Error("failed to establish torrent creation request to the Deluge daemon.")
		return "", err
	}

	if response["result"] == nil {
		c.logger.WithFields(logrus.Fields{"url": c.Url, "err": err}).Error("failed to add torrent.")
		return "", ErrAddTorrentFailed
	}

	return response["result"].(string), nil
}
