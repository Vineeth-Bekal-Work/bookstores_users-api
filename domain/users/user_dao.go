package users

import (
	"fmt"

	"github.com/Vineeth-BekalWork/bookstores_users-api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.Email = result.Email
	user.LastName = result.LastName
	user.DateCreated = result.DateCreated

	return nil
}
func (user *User) Save() *errors.RestErr {
	current := usersDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("Email %d alredy exists", user.Id))
		}
		return errors.NewBadRequestError(fmt.Sprintf("User %d alredy exists", user.Id))
	}
	usersDB[user.Id] = user
	return nil
}
