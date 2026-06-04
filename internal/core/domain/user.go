package domain

import (
	"fmt"
	"regexp"

	core_errors "github.com/Mommsent/todoapp-Studying.git/internal/core/errors"
)

type User struct {
	ID      int
	Version int

	FullName    string
	PhoneNumber *string
}

func NewUser(id int, version int, fullName string, phoneNumber *string) User {
	return User{
		ID:          id,
		Version:     version,
		FullName:    fullName,
		PhoneNumber: phoneNumber,
	}
}

func NewUserUninitialized(fullName string, phoneNumber *string) User {
	return NewUser(
		UninitializedID,
		UnitializedVersion,
		fullName,
		phoneNumber)
}

func (user *User) Validate() error {
	fullNameLenght := len([]rune(user.FullName))

	if fullNameLenght < 3 || fullNameLenght > 100 {
		return fmt.Errorf(
			"invalid `FullName` len: %d: %w", fullNameLenght, core_errors.ErrInvalidArgument)
	}

	if user.PhoneNumber != nil {
		phoneNumberLenght := len([]rune(*user.PhoneNumber))
		if phoneNumberLenght < 3 || phoneNumberLenght > 15 {
			return fmt.Errorf(
				"invalid `PhoneNumber` len: %d: %w", phoneNumberLenght, core_errors.ErrInvalidArgument)
		}

		regular := regexp.MustCompile(`^\+[0-9]+$`)

		if !regular.MatchString(*user.PhoneNumber) {
			return fmt.Errorf("invalid `PhoneNumber` format: %w", core_errors.ErrInvalidArgument)
		}
	}

	return nil
}

type UserPatch struct {
	FullName    Nullable[string]
	PhoneNumber Nullable[string]
}

func (p *UserPatch) Validate() error {
	if p.FullName.Set && p.FullName.Value == nil {
		return fmt.Errorf(
			"'FullName' cant be patched to NULL: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}

func (u *User) ApplyPatch(patch UserPatch) error {
	if err := patch.Validate(); err != nil {
		return fmt.Errorf("validate user patch: %w", err)
	}

	tmp := *u

	if patch.FullName.Set {
		tmp.FullName = *patch.FullName.Value
	}

	if patch.PhoneNumber.Set {
		tmp.PhoneNumber = patch.PhoneNumber.Value
	}

	if err := tmp.Validate(); err != nil {
		return fmt.Errorf("validated pathced user: %w", err)
	}

	*u = tmp

	return nil
}
