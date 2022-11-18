package user

import (
	user_entities_interface "github.com/spro80/golangCleanArchitecture/app/domain/entity/user/interfaces"
)

type userClient struct {
	user *User
}

func NewUserEntity() user_entities_interface.UserEntityInterface {
	userClient := &userClient{
		user: &User{},
	}
	return userClient
}

func (u *userClient) GetRut() string {
	return u.user.Rut
}

func (u *userClient) GetUserName() string {
	return u.user.UserName
}

func (u *userClient) GetPassword() string {
	return u.user.Password
}

func (u *userClient) GetEmail() string {
	return u.user.Email
}

func (u *userClient) GetFirstName() string {
	return u.user.FirstName
}

func (u *userClient) GetLastName() string {
	return u.user.LastName
}

func (u *userClient) SetRut(rut string) error {
	u.user.Rut = rut
	return nil
}

func (u *userClient) SetUserName(name string) error {
	u.user.UserName = name
	return nil
}

func (u *userClient) SetPassword(password string) error {
	u.user.Password = password
	return nil
}

func (u *userClient) SetEmail(email string) error {
	u.user.Email = email
	return nil
}

func (u *userClient) SetFirstName(firstName string) error {
	u.user.FirstName = firstName
	return nil
}

func (u *userClient) SetLastName(lastName string) error {
	u.user.LastName = lastName
	return nil
}
