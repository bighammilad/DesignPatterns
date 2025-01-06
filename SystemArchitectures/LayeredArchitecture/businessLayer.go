package main

import "fmt"

// UserService contains business logic for managing users.
type UserService struct {
	repo *UserRepository
}

// NewUserService creates a new UserService.
func NewUserService(repo *UserRepository) *UserService {
	return &UserService{repo: repo}
}

// CreateUser creates a new user in the system.
func (service *UserService) CreateUser(name string, age int) User {
	// For simplicity, we generate the ID based on the current count.
	user := User{ID: len(service.repo.users) + 1, Name: name, Age: age}
	service.repo.Save(user)
	return user
}

// GetUser retrieves a user by ID.
func (service *UserService) GetUser(id int) (User, error) {
	user, found := service.repo.FindByID(id)
	if !found {
		return User{}, fmt.Errorf("user with ID %d not found", id)
	}
	return user, nil
}

// ListAllUsers returns all users.
func (service *UserService) ListAllUsers() []User {
	return service.repo.GetAllUsers()
}
