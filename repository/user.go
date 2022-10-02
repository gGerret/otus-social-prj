package repository

import (
	"database/sql"
	"errors"
	"github.com/gGerret/otus-social-prj/repository/model"
	"math/rand"
	"time"
)

var ErrorUserNotFound = errors.New("user not found")

type UserRepository struct {
	BaseRepository
}

func GetUserRepository() *UserRepository {
	//Пока без БД балуемся
	/*if database == nil {
	    panic("Error! Database connection is not initialized")
	}*/
	return GetUserRepositoryDB(nil)
}

func GetUserRepositoryDB(db *sql.DB) *UserRepository {
	rep := &UserRepository{}
	rep.db = db
	return rep
}

func (r *UserRepository) GetById(userId string) (*model.UserModel, error) {
	//Пока возвращаем мок
	return CreateUserModelMoc(userId), nil
}

func (r *UserRepository) GetByPublicId(publicId string) (*model.UserModel, error) {
	//Пока возвращаем мок
	return CreateUserModelMoc(), nil
}

func CreateUserModelMoc() *model.UserModel {
	return &model.UserModel{
		Id:         rand.Int63n(2000000),
		PublicId:   userId,
		FirstName:  "Валерий",
		LastName:   "Сидоров",
		MiddleName: "Владимирович",
		Town:       "Сочи",
		Gender:     2,
		CreatedAt:  time.Now().AddDate(0, -1, 0),
		UpdatedAt:  time.Now().AddDate(0, 0, -11),
	}
}
