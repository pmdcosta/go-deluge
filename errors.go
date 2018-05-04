package deluge

// Error represents a Deluge RPC error.
type Error string

// Error returns the error message.
func (e Error) Error() string { return string(e) }

// Deluge RPC errors.
const (
	ErrAuthFailed  = Error("authentication failed")
	ErrInvalidResp = Error("invalid response received")
)
