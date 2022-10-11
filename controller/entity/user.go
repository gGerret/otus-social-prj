package entity

import (
	"github.com/gGerret/otus-social-prj/repository/model"
	"time"
)

type IUserEntity interface {
	FomModel(userModel *model.UserModel)
}

type UserPost struct {
	UserId      string    `json:"user_id"`
	PostId      string    `json:"post_id"`
	PostMessage string    `json:"message"`
	PostTags    []string  `json:"tags"`
	PostDate    time.Time `json:"created"`
}

type UserPublicEntity struct {
	UserId     string   `json:"user_id"`
	FirstName  string   `json:"first_name"`
	LastName   string   `json:"last_name"`
	MiddleName string   `json:"middle_name"`
	Gender     string   `json:"gender"`
	Interests  []string `json:"interests"`
	Town       string   `json:"town"`
}

type UserEntity struct {
	UserId     string    `json:"user_id"`
	Email      string    `json:"email"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	MiddleName string    `json:"middle_name"`
	Gender     string    `json:"gender"`
	Interests  []string  `json:"interests"`
	Town       string    `json:"town"`
	CreatedAt  time.Time `json:"created"`
	UpdatedAt  time.Time `json:"updated"`
}

type UserUpdateEntity struct {
	FirstName  string   `json:"first_name"`
	LastName   string   `json:"last_name"`
	MiddleName string   `json:"middle_name"`
	Gender     string   `json:"gender"`
	Interests  []string `json:"interests"`
	Town       string   `json:"town"`
}

type UserPasswordUpdateEntity struct {
	OldPassword    string `json:"old_pass"`
	NewPassword    string `json:"new_pass"`
	RetypePassword string `json:"retype_pass"`
}

func (u *UserEntity) FromModel(userModel *model.UserModel) {
	u.UserId = userModel.PublicId
	u.Email = userModel.Email
	u.FirstName = userModel.FirstName
	u.LastName = userModel.LastName
	u.MiddleName = userModel.MiddleName
	u.Town = userModel.Town
	u.Gender = userModel.GenderDesc
	u.Interests = userModel.Interests
	u.CreatedAt = userModel.CreatedAt
	u.UpdatedAt = userModel.UpdatedAt
}

func (u *UserEntity) ToModel() *model.UserModel {
	return &model.UserModel{
		PublicId:   u.UserId,
		Email:      u.Email,
		FirstName:  u.FirstName,
		LastName:   u.LastName,
		MiddleName: u.MiddleName,
		Town:       u.Town,
		GenderDesc: u.Gender,
		Interests:  u.Interests,
	}
}

func (u *UserPublicEntity) FromModel(userModel *model.UserModel) {
	u.UserId = userModel.PublicId
	u.FirstName = userModel.FirstName
	u.LastName = userModel.LastName
	u.MiddleName = userModel.MiddleName
	u.Town = userModel.Town
	u.Gender = userModel.GenderDesc
	u.Interests = userModel.Interests
}

func (u *UserUpdateEntity) FromModel(userModel *model.UserModel) {
	u.FirstName = userModel.FirstName
	u.LastName = userModel.LastName
	u.MiddleName = userModel.MiddleName
	u.Town = userModel.Town
	u.Gender = userModel.GenderDesc
	u.Interests = userModel.Interests
}

func (u *UserUpdateEntity) ToModel() *model.UserModel {
	return &model.UserModel{
		FirstName:  u.FirstName,
		LastName:   u.LastName,
		MiddleName: u.MiddleName,
		Town:       u.Town,
		GenderDesc: u.Gender,
		Interests:  u.Interests,
	}
}
