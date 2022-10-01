package api

import (
	"errors"
	"strings"
)

type UserService interface {
	New(user User) error
	GetAll() ([]*User, error)
	GetOne(id int) (*User, error)
	Update(user *User) error
	Delete(id int) error
}

type UserRepository interface {
	CreateUser(User) error
	GetAllUsers() ([]*User, error)
	GetOneUser(int) (*User, error)
	UpdateUser(*User) error
	DeleteUser(int) error
}

type userService struct {
	storage UserRepository
}

func NewUserService(userRepo UserRepository) UserService {
	return &userService{
		storage: userRepo,
	}
}

func (u *userService) New(user User) error {
	if user.Email == "" {
		return errors.New("email is required")
	}

	if user.Password == "" {
		return errors.New("password is required")
	}

	user.Email = strings.ToLower(user.Email)

	err := u.storage.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (u *userService) GetAll() ([]*User, error) {
	users, err := u.storage.GetAllUsers()

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *userService) GetOne(id int) (*User, error) {
	post, err := u.storage.GetOneUser(id)

	if err != nil {
		return nil, err
	}

	return post, nil
}

func (u *userService) Update(user *User) error {
	user.Email = strings.ToLower(user.Email)
	user.Username = strings.ToLower(user.Username)
	user.FirstName = strings.ToLower(user.FirstName)
	user.LastName = strings.ToLower(user.LastName)

	err := u.storage.UpdateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (u *userService) Delete(id int) error {
	err := u.storage.DeleteUser(id)

	if err != nil {
		return err
	}

	return nil
}
