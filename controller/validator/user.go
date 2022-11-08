package validator

import (
	"github.com/gGerret/otus-social-prj/controller/entity"
	"regexp"
)

var (
	UserIdFieldError   = FieldError{"user_id", "UserId format in incorrect"}
	EmailFieldErr      = FieldError{"email", "Email format is incorrect"}
	FirstNameFieldErr  = FieldError{"first_name", "First name is incorrect"}
	LastNameFieldErr   = FieldError{"last_name", "First name is incorrect"}
	MiddleNameFieldErr = FieldError{"middle_name", "First name is incorrect"}
	PasswdFieldErr     = FieldError{"passwd", "Password must contains from 6 to 20 symbols. At leas one digit, one lowercase latin letter and one uppercase latin letter"}
	PasswdNotMatch     = FieldError{"passwd", "Password and retype does not match"}
	InterestInvalid    = FieldError{"interests", "Interest string is not valid"}
)

type UserRegisterValidator struct {
	Entity *entity.UserRegisterEntity
}

func (v *UserRegisterValidator) Validate() []*FieldError {
	var fieldErrs []*FieldError

	if !isValidEmail(v.Entity.Email) {
		fieldErrs = append(fieldErrs, &EmailFieldErr)
	}

	if !isValidName(v.Entity.FirstName) {
		fieldErrs = append(fieldErrs, &FirstNameFieldErr)
	}
	if !isValidName(v.Entity.LastName) {
		fieldErrs = append(fieldErrs, &LastNameFieldErr)
	}
	if !isValidName(v.Entity.MiddleName) {
		fieldErrs = append(fieldErrs, &MiddleNameFieldErr)
	}
	if !isValidPassword(v.Entity.Password) {
		fieldErrs = append(fieldErrs, &PasswdFieldErr)
	}
	if !isPasswordMatchRetype(v.Entity.Password, v.Entity.RetypePassword) {
		fieldErrs = append(fieldErrs, &PasswdNotMatch)
	}

	return fieldErrs
}

type NewFriendValidator struct {
	Entity *entity.NewFriendPublicIdEntity
}

func (v *NewFriendValidator) Validate() []*FieldError {
	var fieldErrs []*FieldError

	if !isValidUserId(v.Entity.UserId) {
		fieldErrs = append(fieldErrs, &UserIdFieldError)
	}

	return fieldErrs
}

type UserFilterValidator struct {
	Entity *entity.UserFilterEntity
}

func (v *UserFilterValidator) Validate() []*FieldError {
	var fieldErrs []*FieldError

	if !isValidName(v.Entity.FirstName) {
		fieldErrs = append(fieldErrs, &FirstNameFieldErr)
	}
	if !isValidName(v.Entity.LastName) {
		fieldErrs = append(fieldErrs, &LastNameFieldErr)
	}
	if !isValidName(v.Entity.MiddleName) {
		fieldErrs = append(fieldErrs, &MiddleNameFieldErr)
	}
	for _, interest := range v.Entity.Interests {
		if !isValidInterest(interest) {
			fieldErrs = append(fieldErrs, &InterestInvalid)
			break
		}
	}
	return fieldErrs
}

func isValidEmail(email string) bool {
	var validEmail = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,7}$`)
	return validEmail.MatchString(email)
}
func isValidPassword(password string) bool {
	var validPassword = regexp.MustCompile(`(\d*)([a-z]*)([A-Z]*).{6,20}`)
	return validPassword.MatchString(password)
}
func isPasswordMatchRetype(password string, retype string) bool {
	return password == retype
}
func isValidUserId(userId string) bool {
	var validUserId = regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)
	return validUserId.MatchString(userId)
}
func isValidName(someName string) bool {
	if len(someName) == 0 {
		return true
	}
	var validName = regexp.MustCompile(`^[a-zA-Zа-яА-Я'\s]+$`)
	return validName.MatchString(someName)
}
func isValidInterest(someInterest string) bool {
	if len(someInterest) == 0 {
		return false
	}
	var validInterest = regexp.MustCompile(`^[a-zA-Zа-яА-Я0-9'_\-\s"]+$`)
	return validInterest.MatchString(someInterest)
}
