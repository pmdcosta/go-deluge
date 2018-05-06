package deluge

// service represents a service for interacting with deluge.
type Service struct {
	client *Client
}

// AddTorrentMagnet starts downloading a magnet link.
// magnetUrl is the Magnet URL for the torrent.
// options is a map with options to be set. (check the Deluge Torrent documentation).
func (s *Service) AddTorrentMagnet(magnetUrl string, options map[string]interface{}) (string, error) {
	if magnetUrl == "" {
		return "", ErrInvalidTorrent
	}

	response, err := s.client.sendRequest(ADD_MAGNET, magnetUrl, options)
	if err != nil {
		return "", err
	}

	if response["result"] == nil {
		return "", ErrInvalidResp
	}

	return response["result"].(string), nil
}

// AddTorrentUrl starts downloading a torrent from a URL.
// torrentUrl is the URL for the torrent.
// options is a map with options to be set. (check the Deluge Torrent documentation).
func (s *Service) AddTorrentUrl(torrentUrl string, options map[string]interface{}) (string, error) {
	if torrentUrl == "" {
		return "", ErrInvalidTorrent
	}

	response, err := s.client.sendRequest(ADD_TORRENT_URL, torrentUrl, options)
	if err != nil {
		return "", err
	}

	if response["result"] == nil {
		return "", ErrInvalidResp
	}

	return response["result"].(string), nil
}

// AddTorrentFile starts downloading a torrent through a file.
// fileName is the name of the original torrent file.
// fileDump is the base64 encoded contents of the file.
// options is a map with options to be set. (check the Deluge Torrent documentation).
func (s *Service) AddTorrentFile(fileName, fileDump string, options map[string]interface{}) (string, error) {
	response, err := s.client.sendRequest(ADD_TORRENT_FILE, fileName, fileDump, options)
	if err != nil {
		return "", err
	}

	if response["result"] == nil {
		return "", ErrInvalidResp
	}

	return response["result"].(string), nil
}
