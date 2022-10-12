package validator

import (
	"github.com/gGerret/otus-social-prj/controller/entity"
	"regexp"
)

var (
	EmailFieldErr      = FieldError{"email", "Email format is incorrect"}
	FirstNameFieldErr  = FieldError{"first_name", "First name is incorrect"}
	LastNameFieldErr   = FieldError{"last_name", "First name is incorrect"}
	MiddleNameFieldErr = FieldError{"middle_name", "First name is incorrect"}
	PasswdFieldErr     = FieldError{"passwd", "Password must contains from 6 to 20 symbols. At leas one digit, one lowercase latin letter and one uppercase latin letter"}
	PasswdNotMatch     = FieldError{"passwd", "Password and retype does not match"}
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

func isValidEmail(email string) bool {
	var validEmail = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,7}$`)
	return validEmail.MatchString(email)
}
func isValidPassword(password string) bool {
	var validPassword = regexp.MustCompile(`(?=\d*)(?=[a-z]*)(?=[A-Z]*).{6,20}`)
	return validPassword.MatchString(password)
}
func isPasswordMatchRetype(password string, retype string) bool {
	return password == retype
}
func isValidName(someName string) bool {
	var validName = regexp.MustCompile(`^[a-zA-Zа-яА-Я]$`)
	return validName.MatchString(someName)
}
