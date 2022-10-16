package repository

import (
	"database/sql"
	"errors"
	"fmt"
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
	if database == nil {
		panic("Error! Database connection is not initialized")
	}
	return GetUserRepositoryDB(database)
}

func GetUserRepositoryDB(db *sql.DB) *UserRepository {
	rep := &UserRepository{}
	rep.db = db
	return rep
}

func (r *UserRepository) GetById(userId int64) (*model.UserModel, error) {
	row := r.db.QueryRow(
		"select usr.id,\n"+
			"       usr.public_id,\n"+
			"       usr.pass_hash,\n"+
			"       usr.email,\n"+
			"       usr.first_name,\n"+
			"       usr.last_name,\n"+
			"       usr.middle_name,\n"+
			"       usr.gender AS gender,\n"+
			"       g.full_desc AS gender_desc,\n"+
			"       usr.town,\n"+
			"       usr.created_at,\n"+
			"       usr.updated_at,\n"+
			"       usr.deleted_at\n"+
			"from social.user usr\n"+
			"    join social.gender g ON g.id = usr.gender\n"+
			"where usr.id = ?", userId)
	userModel := model.UserModel{}
	if err := row.Scan(&userModel.Id, &userModel.PublicId, &userModel.PasswordHash, &userModel.Email,
		&userModel.FirstName, &userModel.LastName, &userModel.MiddleName, &userModel.Gender, &userModel.GenderDesc, &userModel.Town,
		&userModel.CreatedAt, &userModel.UpdatedAt, &userModel.DeletedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("UserRepository.GetById %d : no such user", userId)
		}
		return nil, fmt.Errorf("InterestRepository.GetById %d: %v", userId, err)
	}
	return &userModel, nil
}

func (r *UserRepository) GetByPublicId(publicId string) (*model.UserModel, error) {
	row := r.db.QueryRow(
		"select usr.id,\n"+
			"       usr.public_id,\n"+
			"       usr.pass_hash,\n"+
			"       usr.email,\n"+
			"       usr.first_name,\n"+
			"       usr.last_name,\n"+
			"       usr.middle_name,\n"+
			"       usr.gender AS gender,\n"+
			"       g.full_desc AS gender_desc,\n"+
			"       usr.town,\n"+
			"       usr.created_at,\n"+
			"       usr.updated_at,\n"+
			"       usr.deleted_at\n"+
			"from social.user usr\n"+
			"    join social.gender g ON g.id = usr.gender\n"+
			"where usr.public_id = ?", &publicId)
	userModel := model.UserModel{}
	if err := row.Scan(&userModel.Id, &userModel.PublicId, &userModel.PasswordHash, &userModel.Email,
		&userModel.FirstName, &userModel.LastName, &userModel.MiddleName, &userModel.Gender, &userModel.GenderDesc, &userModel.Town,
		&userModel.CreatedAt, &userModel.UpdatedAt, &userModel.DeletedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("UserRepository.GetByPublicId %s : no such user", publicId)
		}
		return nil, fmt.Errorf("InterestRepository.GetByPublicId %s: %v", publicId, err)
	}
	return &userModel, nil
}

func (r *UserRepository) UpdateUser(userModel *model.UserModel) error {
	return fmt.Errorf("UserRepository.UpdateUser %d, %s: method is not implemented yet", userModel.Id, userModel.PublicId)
}

//TODO: Добавить транзакцию для атомарного добавления пользователя
func (r *UserRepository) CreateByModel(userModel *model.UserModel) (createdUser *model.UserModel, err error) {

	usrRawModel := model.GetRawUserModel(userModel)
	interestRepo := GetInterestRepositoryDB(r.db)

	knownInterests, unknownInterests, err := interestRepo.GetInterestsFromList(userModel.Interests)
	if err != nil {
		return nil, err
	}

	if len(unknownInterests) > 0 {
		newInterests, err := interestRepo.InsertInterestsSkipExisting(unknownInterests)
		if err != nil {
			return nil, err
		}
		knownInterests = append(knownInterests, newInterests...)
	}

	result, err := r.db.Exec("insert into social.user (public_id, pass_hash, email, first_name, last_name, middle_name, gender, town, created_at) "+
		"values (?, ?, ?, ?, ?, ?, ?, ?, ?)",
		usrRawModel.PublicId, usrRawModel.PasswordHash, usrRawModel.Email, usrRawModel.FirstName,
		usrRawModel.LastName, usrRawModel.MiddleName, usrRawModel.Gender, usrRawModel.Town, usrRawModel.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	newUser := *userModel
	newUser.Id, err = result.LastInsertId()
	if err != nil {
		return nil, err
	}
	createdUser = &newUser
	if knownInterests != nil {
		err = interestRepo.LinkInterestsToUser(createdUser.Id, knownInterests)
		if err != nil {
			return createdUser, err
		}
	}

	return createdUser, nil
}

//Функция для полного удаления пользователя. В первую очередь для тестов,
//но возможно нужно будет использовать для исполнения закона о цифровом забвении № 149-ФЗ
func (r *UserRepository) ForceUserDelete(userModel *model.UserModel) error {
	interestRepo := GetInterestRepositoryDB(r.db)
	err := interestRepo.UnlinkAllUserInterests(userModel.Id)
	if err != nil {
		return err
	}
	_, err = r.db.Exec("delete from social.user where id = ?", userModel.Id)
	return err
}

func CreateUserModelMoc() *model.UserModel {
	newId := rand.Int63n(2000000)
	return &model.UserModel{
		Id:         newId,
		PublicId:   uuid.Must(uuid.NewV4()).String(),
		Email:      fmt.Sprintf("sidorov_%d@yandex.ru", newId),
		FirstName:  "Валерий",
		LastName:   "Сидоров",
		MiddleName: "Владимирович",
		Town:       "Сочи",
		Gender:     2,
		GenderDesc: "Мужской",
		Interests:  []string{"Автомобили", "Рисование", "Программирование", "Вышивка"},
		CreatedAt:  sql.NullTime{Time: time.Now().AddDate(0, -1, 0), Valid: true},
		UpdatedAt:  sql.NullTime{Time: time.Now().AddDate(0, 0, -11), Valid: true},
	}
}
