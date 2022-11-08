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
	Gender     int      `json:"gender"`
	GenderDesc string   `json:"gender_desc"`
	Interests  []string `json:"interests"`
	Town       string   `json:"town"`
}

type UserPublicEntityArray struct {
	Users []UserPublicEntity `json:"users"`
}

type UserEntity struct {
	UserId     string    `json:"user_id"`
	Email      string    `json:"email"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	MiddleName string    `json:"middle_name"`
	Gender     int       `json:"gender"`
	GenderDesc string    `json:"gender_desc"`
	Interests  []string  `json:"interests"`
	Town       string    `json:"town"`
	CreatedAt  time.Time `json:"created"`
	UpdatedAt  time.Time `json:"updated"`
}

type UserRegisterEntity struct {
	Email          string   `json:"email"`
	Password       string   `json:"passwd"`
	RetypePassword string   `json:"retype_passwd"`
	FirstName      string   `json:"first_name"`
	LastName       string   `json:"last_name"`
	MiddleName     string   `json:"middle_name"`
	Gender         int      `json:"gender"`
	Interests      []string `json:"interests"`
	Town           string   `json:"town"`
}

type UserUpdateEntity struct {
	FirstName  string   `json:"first_name"`
	LastName   string   `json:"last_name"`
	MiddleName string   `json:"middle_name"`
	Gender     int      `json:"gender"`
	Interests  []string `json:"interests"`
	Town       string   `json:"town"`
}

type UserFilterEntity struct {
	FirstName  string   `json:"first_name"`
	LastName   string   `json:"last_name"`
	MiddleName string   `json:"middle_name"`
	Gender     int      `json:"gender"`
	Interests  []string `json:"interests"`
	Town       string   `json:"town"`
}

type UserPasswordUpdateEntity struct {
	OldPassword    string `json:"old_pass"`
	NewPassword    string `json:"new_pass"`
	RetypePassword string `json:"retype_pass"`
}

type UserBaseLoginEntity struct {
	Email    string `json:"email"`
	Password string `json:"passwd"`
}

type NewFriendPublicIdEntity struct {
	UserId  string `json:"user_id"`
	Comment string `json:"comment"`
}

func (u *UserEntity) FromModel(userModel *model.UserModel) {
	u.UserId = userModel.PublicId
	u.Email = userModel.Email
	u.FirstName = userModel.FirstName
	u.LastName = userModel.LastName
	u.MiddleName = userModel.MiddleName
	u.Town = userModel.Town
	u.Gender = userModel.Gender
	u.GenderDesc = userModel.GenderDesc
	u.Interests = userModel.Interests
	u.CreatedAt = userModel.CreatedAt.Time
	u.UpdatedAt = userModel.UpdatedAt.Time
}

func (u *UserEntity) ToModel() *model.UserModel {
	return &model.UserModel{
		Id:         0,
		PublicId:   u.UserId,
		Email:      u.Email,
		FirstName:  u.FirstName,
		LastName:   u.LastName,
		MiddleName: u.MiddleName,
		Town:       u.Town,
		Gender:     u.Gender,
		GenderDesc: u.GenderDesc,
		Interests:  u.Interests,
	}
}

func (u *UserPublicEntity) FromModel(userModel *model.UserModel) {
	u.UserId = userModel.PublicId
	u.FirstName = userModel.FirstName
	u.LastName = userModel.LastName
	u.MiddleName = userModel.MiddleName
	u.Town = userModel.Town
	u.Gender = userModel.Gender
	u.GenderDesc = userModel.GenderDesc
	u.Interests = userModel.Interests
}

func (u *UserUpdateEntity) FromModel(userModel *model.UserModel) {
	u.FirstName = userModel.FirstName
	u.LastName = userModel.LastName
	u.MiddleName = userModel.MiddleName
	u.Town = userModel.Town
	u.Gender = userModel.Gender
	u.Interests = userModel.Interests
}

func (u *UserPublicEntity) ToModel() *model.UserModel {
	return &model.UserModel{
		Id:         0,
		PublicId:   u.UserId,
		FirstName:  u.FirstName,
		LastName:   u.LastName,
		MiddleName: u.MiddleName,
		Town:       u.Town,
		Gender:     u.Gender,
		GenderDesc: u.GenderDesc,
		Interests:  u.Interests,
	}
}

func (u *UserUpdateEntity) ToModel() *model.UserModel {
	return &model.UserModel{
		Id:         0,
		FirstName:  u.FirstName,
		LastName:   u.LastName,
		MiddleName: u.MiddleName,
		Town:       u.Town,
		Gender:     u.Gender,
		Interests:  u.Interests,
	}
}
func (u *UserRegisterEntity) ToModel() *model.UserModel {
	return &model.UserModel{
		Id:         0,
		Email:      u.Email,
		FirstName:  u.FirstName,
		LastName:   u.LastName,
		MiddleName: u.MiddleName,
		Town:       u.Town,
		Gender:     u.Gender,
		Interests:  u.Interests,
	}
}

func (u *NewFriendPublicIdEntity) ToModel() *model.UserModel {
	return &model.UserModel{
		Id:       0,
		PublicId: u.UserId,
	}
}
func (g *UserPublicEntityArray) FromModelArray(ua []model.UserModel) {
	for _, userModel := range ua {
		userEntity := UserPublicEntity{}
		userEntity.FromModel(&userModel)
		g.Users = append(g.Users, userEntity)
	}
}

func (f *UserFilterEntity) ToModel() *model.UserFilterModel {
	m := &model.UserFilterModel{}
	if len(f.FirstName) > 0 {
		m.FirstName = &f.FirstName
	}
	if len(f.LastName) > 0 {
		m.LastName = &f.LastName
	}
	if len(f.MiddleName) > 0 {
		m.MiddleName = &f.MiddleName
	}
	if f.Gender != 0 {
		m.Gender = &f.Gender
	}
	if len(f.Town) > 0 {
		m.Town = &f.Town
	}
	m.Interests = f.Interests

	return m
}
