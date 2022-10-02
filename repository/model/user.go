package model

import "time"

type UserModel struct {
	Id           int64
	PublicId     string
	PasswordHash string
	FirstName    string
	LastName     string
	MiddleName   string
	Gender       int
	GenderDesc   string
	Interests    []string
	Town         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}

func (u *UserModel) TableName() string {
	return "user"
}
