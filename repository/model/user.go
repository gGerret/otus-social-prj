package model

import "time"

type UserModel struct {
	Id           int64
	PublicId     string
	Email        string
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

type UserRawModel struct {
	Id           int64
	PublicId     string
	Email        string
	PasswordHash string
	FirstName    string
	LastName     string
	MiddleName   string
	Gender       int
	Town         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}

func GetRawUserModel(um *UserModel) *UserRawModel {
	return &UserRawModel{
		Id:           um.Id,
		PublicId:     um.PublicId,
		Email:        um.Email,
		PasswordHash: um.PasswordHash,
		FirstName:    um.FirstName,
		LastName:     um.LastName,
		MiddleName:   um.MiddleName,
		Gender:       um.Gender,
		Town:         um.Town,
		CreatedAt:    um.CreatedAt,
		UpdatedAt:    um.UpdatedAt,
		DeletedAt:    um.DeletedAt,
	}
}

func (u *UserRawModel) TableName() string {
	return "user"
}
