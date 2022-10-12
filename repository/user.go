package repository

import (
	"database/sql"
	"errors"
	"github.com/gGerret/otus-social-prj/repository/model"
	"github.com/gofrs/uuid"
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
	return CreateUserModelMoc(), nil
}

func (r *UserRepository) GetByPublicId(publicId string) (*model.UserModel, error) {
	//Пока возвращаем мок
	return CreateUserModelMoc(), nil
}

func (r *UserRepository) UpdateUser(userModel *model.UserModel) error {
	return nil
}

func (r *UserRepository) CreateByModel(userModel *model.UserModel) error {

	usrRawModel := model.GetRawUserModel(userModel)
	interestRepo := GetInterestRepositoryDB(r.db)

	userInterests := userModel.Interests

	knownInterests, err := interestRepo.GetKnownInterestsFromList(userModel.Interests)
	for _,interest := userModel.Interests {

	}

	createUsrTx, err := r.db.Begin()
	if err != nil {
		return err
	}

	_, err = createUsrTx.Exec("insert into social.user (public_id, pass_hash, email, first_name, last_name, middle_name, gender, town, created_at) " +
		"values (?, ?, ?, ?, ?, ?, ?, ?, ?)",
		usrRawModel.PublicId, usrRawModel.PasswordHash, usrRawModel.Email, usrRawModel.FirstName,
		usrRawModel.LastName, usrRawModel.FirstName, usrRawModel.MiddleName, usrRawModel.Gender,
		usrRawModel.Town, usrRawModel.CreatedAt,
	)
	if err != nil {
		return err
	}


	return nil
}

func CreateUserModelMoc() *model.UserModel {
	return &model.UserModel{
		Id:         rand.Int63n(2000000),
		PublicId:   uuid.Must(uuid.NewV4()).String(),
		Email:      "sidorov@yandex.ru",
		FirstName:  "Валерий",
		LastName:   "Сидоров",
		MiddleName: "Владимирович",
		Town:       "Сочи",
		Gender:     2,
		GenderDesc: "Мужской",
		Interests:  []string{"Автомобили", "Рисование", "Программирование"},
		CreatedAt:  time.Now().AddDate(0, -1, 0),
		UpdatedAt:  time.Now().AddDate(0, 0, -11),
	}
}
