package deluge

// Error represents a Deluge RPC error
type Error string

// Error returns the error message
func (e Error) Error() string { return string(e) }

// Deluge RPC errors
const (
	ErrAuthFailed       = Error("authentication failed")
	ErrRequestFailed    = Error("deluge request failed")
	ErrAddTorrentFailed = Error("failed to add torrent")
	ErrInvalidTorrent   = Error("invalid torrent url")
)
