package services

import "fmt"

// UserService user service type
type UserService struct {
	CreateUser func()
	GetUser    func()
}

// NewUserService creates a new service layer
func NewUserService() UserService {
	return UserService{
		CreateUser: createUser,
		GetUser:    getUser,
	}
}

func createUser() {
	fmt.Println("FROM SERVICE: Create User Called.")
}

func getUser() {
	fmt.Println("FROM SERVICE: Get user service called.")
}
