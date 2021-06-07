package server

// Abstract Session Repository.
type ABSSessionsRepository interface {
	Add(ses Session)  // Add a Session to repository
	Remove(id uint64) bool  // Remove an item from the repository.
	Get(id uint64) (Session, bool) // Get a Session by id.
	GetAll()  // Return the slice of elements inside the repository.
	Size() int  // Return the size of the Repository.
}

// In memory implementation of the Session Repository.
type SessionsRepository struct {
	elements []Session
}

// Create a new SessionsRepository.
func NewSessionsRepository() *SessionsRepository {
	return &SessionsRepository{elements: nil}
}

// Add a Session to the Repository.
func (s *SessionsRepository) Add(ses Session) {
	s.elements = append(s.elements, ses)
}

// Remove Session with given id.
// Returns true if item with given id was found.
// Returns false if that item could not be found.
func (s *SessionsRepository) Remove(id uint64) bool {
	panic("implement me")
}

// Return a slice of Session inside the Repository.
func (s *SessionsRepository) GetAll() {
	panic("implement me")
}

// Get the number of Session inside the Repository.
func (s *SessionsRepository) Size() int {
	return len(s.elements)
}

// Return the Session with the given id.
// Returns true if that Session was found.
// If false is returned, the Session will be zero-initialized
func (s *SessionsRepository) Get(id uint64) (Session, bool) {

}
