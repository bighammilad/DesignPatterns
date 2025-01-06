package main

// User represents a user in the system.
type User struct {
	ID   int
	Name string
	Age  int
}

// UserRepository is responsible for data access and storage (simulated).
type UserRepository struct {
	users map[int]User
}

// NewUserRepository creates a new user repository.
func NewUserRepository() *UserRepository {
	return &UserRepository{users: make(map[int]User)}
}

// Save saves a user to the repository.
func (repo *UserRepository) Save(user User) {
	repo.users[user.ID] = user
}

// FindByID retrieves a user by their ID.
func (repo *UserRepository) FindByID(id int) (User, bool) {
	user, exists := repo.users[id]
	return user, exists
}

// GetAllUsers returns all users.
func (repo *UserRepository) GetAllUsers() []User {
	var allUsers []User
	for _, user := range repo.users {
		allUsers = append(allUsers, user)
	}
	return allUsers
}
