package deluge

// Method represents a deluge RPC request type.
type Method string

// Deluge methods.
const (
	AUTH             = Method("auth.login")
	ADD_MAGNET       = Method("core.add_torrent_magnet")
	ADD_TORRENT_URL  = Method("core.add_torrent_url")
	ADD_TORRENT_FILE = Method("core.add_torrent_file")
)
