package server

// Stores session data for the user.
type Session struct {
	id uint64  // Session id

}

// Create a new session
func NewSession(id uint64) *Session {
	return &Session{id: id}
}
