package domain

import (
	"fmt"
	"regexp"

	core_errors "github.com/MrLaplace19/ToDoList/internal/core/errors"
)


type User struct{
	ID int 
	Version int

	FullName string
	PhoneNumber *string
}

func NewUserUnilitialized(fullname string, phonenumber *string) User{
	return NewUser(fullname, phonenumber, UninitializedID, UninitiliazideVersion)
}

func NewUser(fullname string, phonenumber *string, id, version int) User{
	return User{
		ID: id,
		Version: version,

		FullName: fullname,
		PhoneNumber: phonenumber,
	}
}

func (u *User) Validate() error{
	fullNameLenght := len([]rune(u.FullName))
	if fullNameLenght < 3 || fullNameLenght > 100 {
		return fmt.Errorf("Invalid `FullName` len %d: %w", fullNameLenght, core_errors.ErrInvalidArgument)
	}

	if u.PhoneNumber != nil{
		phoneNumberLen := len([]rune(*u.PhoneNumber))
		if phoneNumberLen < 10 || phoneNumberLen > 15{
			return fmt.Errorf("invalid `PhoneNumber` len: %d: %w", phoneNumberLen, core_errors.ErrInvalidArgument)

		}

		re := regexp.MustCompile(`^\+[0-9]+$`)

		if !re.MatchString(*u.PhoneNumber){
			return fmt.Errorf("Invalid `PhoneNumber` format: %w ", core_errors.ErrInvalidArgument)
		}

	}

	return nil
}