package deluge

// Method represents a deluge RPC request type.
type Method string

// Deluge methods.
const (
	MethodAuth          = Method("auth.login")
	MethodAddMagnet     = Method("core.add_torrent_magnet")
	MethodAddTorrentURL = Method("core.add_torrent_url")
)
