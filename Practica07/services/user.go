package services

import "fmt"

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

type UserService struct {
    users []User
}

func NewUserService() *UserService {
    return &UserService{
        users: []User{},
    }
}

func (s *UserService) GetUsers() []User {
    return s.users
}

func (s *UserService) CreateUser(user User) User {
    user.ID = len(s.users) + 1
    s.users = append(s.users, user)
    return user
}

func (s *UserService) UpdateUser(id int, updatedUser User) (User, error) {
    for i, user := range s.users {
        if user.ID == id {
            s.users[i].Name = updatedUser.Name
            s.users[i].Email = updatedUser.Email
            return s.users[i], nil
        }
    }
    return User{}, fmt.Errorf("user not found")

}

func (s *UserService) DeleteUser(id int) error {
	for i, user := range s.users {
		if user.ID == id {
			s.users = append(s.users[:i], s.users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("user not found")
}
