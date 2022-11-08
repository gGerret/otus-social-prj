package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/gGerret/otus-social-prj/repository/model"
	"github.com/gofrs/uuid"
	"math/rand"
	"strconv"
	"strings"
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
		return nil, fmt.Errorf("UserRepository.GetById %d: %v", userId, err)
	}
	intRep := GetInterestRepository()
	interests, err := intRep.GetByUserId(userModel.Id)
	if err != nil {
		return nil, fmt.Errorf("UserRepository.GetById %s: %v", userId, err)
	}
	userModel.Interests = interests.Interests
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
		return nil, fmt.Errorf("UserRepository.GetByPublicId %s: %v", publicId, err)
	}

	intRep := GetInterestRepository()
	interests, err := intRep.GetByUserId(userModel.Id)
	if err != nil {
		return nil, fmt.Errorf("UserRepository.GetByPublicId %s: %v", publicId, err)
	}
	userModel.Interests = interests.Interests
	return &userModel, nil
}

func (r *UserRepository) UpdateUser(userModel *model.UserModel) error {
	return fmt.Errorf("UserRepository.UpdateUser %d, %s: method is not implemented yet", userModel.Id, userModel.PublicId)
}

// TODO: Добавить транзакцию для атомарного добавления пользователя
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
		"values (?, ?, ?, ?, ?, ?, ?, ?, now())",
		usrRawModel.PublicId, usrRawModel.PasswordHash, usrRawModel.Email, usrRawModel.FirstName,
		usrRawModel.LastName, usrRawModel.MiddleName, usrRawModel.Gender, usrRawModel.Town,
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

func (r *UserRepository) GetUserFriends(userModel *model.UserModel) (friends []model.UserModel, err error) {
	intRep := GetInterestRepositoryDB(r.db)
	query :=
		"select usr.id,\n" +
			"       usr.public_id,\n" +
			"       usr.pass_hash,\n" +
			"       usr.email,\n" +
			"       usr.first_name,\n" +
			"       usr.last_name,\n" +
			"       usr.middle_name,\n" +
			"       usr.gender AS gender,\n" +
			"       g.full_desc AS gender_desc,\n" +
			"       usr.town,\n" +
			"       usr.created_at,\n" +
			"       usr.updated_at,\n" +
			"       usr.deleted_at\n" +
			"from social.user usr\n" +
			"    join social.user_friendship_link ufl on usr.id = ufl.user_id_b\n" +
			"    join social.gender g ON g.id = usr.gender\n"
	if userModel.Id == 0 {
		query += fmt.Sprintf("where ufl.user_id_a = (select u.id from social.user u where u.public_id = %s limit 1)", userModel.PublicId)
	} else {
		query += fmt.Sprintf("where ufl.user_id_a = %d", userModel.Id)
	}
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("UserRepository.GetUserFriends %d: %v", userModel.Id, err)
	}
	for rows.Next() {
		var friend model.UserModel
		if err := rows.Scan(&friend.Id, &friend.PublicId, &friend.PasswordHash, &friend.Email,
			&friend.FirstName, &friend.LastName, &friend.MiddleName, &friend.Gender, &friend.GenderDesc, &friend.Town,
			&friend.CreatedAt, &friend.UpdatedAt, &friend.DeletedAt); err != nil {
			return nil, fmt.Errorf("UserRepository.GetUserFriends %d: %v", userModel.Id, err)
		}

		interests, err := intRep.GetByUserId(friend.Id)
		if err != nil {
			return nil, fmt.Errorf("UserRepository.GetUserFriends %s: %v", userModel.Id, err)
		}
		friend.Interests = interests.Interests
		friends = append(friends, friend)
	}

	return friends, nil
}

func (r *UserRepository) GetUsersByFilter(filterModel *model.UserFilterModel) (friends []model.UserModel, err error) {
	intRep := GetInterestRepositoryDB(r.db)
	query :=
		"select usr.id,\n" +
			"       usr.public_id,\n" +
			"       usr.pass_hash,\n" +
			"       usr.email,\n" +
			"       usr.first_name,\n" +
			"       usr.last_name,\n" +
			"       usr.middle_name,\n" +
			"       usr.gender AS gender,\n" +
			"       g.full_desc AS gender_desc,\n" +
			"       usr.town,\n" +
			"       usr.created_at,\n" +
			"       usr.updated_at,\n" +
			"       usr.deleted_at\n" +
			"from social.user usr\n" +
			"    join social.gender g ON g.id = usr.gender\n"
	if len(filterModel.Interests) > 0 {
		query += "    join social.user_interests_link il ON il.user_id = usr.id\n" +
			"    join social.interests itr ON il.interest_id = itr.id\n"
	}

	if clause := buildFilterClause(filterModel); len(clause) == 0 {
		return nil, fmt.Errorf("UserRepository.GetUsersByFilter: Empty filter is not permited")
	} else {
		query += clause
	}

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("UserRepository.GetUsersByFilter: %v", err)
	}
	for rows.Next() {
		var friend model.UserModel
		if err := rows.Scan(&friend.Id, &friend.PublicId, &friend.PasswordHash, &friend.Email,
			&friend.FirstName, &friend.LastName, &friend.MiddleName, &friend.Gender, &friend.GenderDesc, &friend.Town,
			&friend.CreatedAt, &friend.UpdatedAt, &friend.DeletedAt); err != nil {
			return nil, fmt.Errorf("UserRepository.GetUsersByFilter: %v", err)
		}

		interests, err := intRep.GetByUserId(friend.Id)
		if err != nil {
			return nil, fmt.Errorf("UserRepository.GetUsersByFilter: %v", err)
		}
		friend.Interests = interests.Interests
		friends = append(friends, friend)
	}

	return friends, nil
}

func buildFilterClause(filterModel *model.UserFilterModel) string {
	clause := "where 1 = 1\n"

	if filterModel.FirstName != nil {
		clause += "    and usr.first_name like '" + *filterModel.FirstName + "'\n"
	}
	if filterModel.LastName != nil {
		clause += "    and usr.last_name like '" + *filterModel.LastName + "'\n"
	}
	if filterModel.MiddleName != nil {
		clause += "    and usr.middle_name like '" + *filterModel.MiddleName + "'\n"
	}
	if filterModel.Town != nil {
		clause += "    and usr.town like '" + *filterModel.Town + "'\n"
	}
	if filterModel.Gender != nil {
		clause += "    and usr.gender = " + strconv.Itoa(*filterModel.Gender) + "\n"
	}
	if len(filterModel.Interests) > 0 {
		clause += "    and itr.interest in ('" + filterModel.Interests[0] + "'"
		for _, interest := range filterModel.Interests[1:] {
			clause += ",'" + interest + "'"
		}
		clause += ")\n"
	}
	return clause
}

func (r *UserRepository) CreateFriendshipLink(userA *model.UserModel, userB *model.UserModel, comment string) (err error) {

	if userA == nil || userA.Id == 0 {
		return fmt.Errorf("UserRepository.CreateFriendshipLink: userA can not be nil")
	}
	if userB == nil {
		return fmt.Errorf("UserRepository.CreateFriendshipLink: userB can not be nil")
	}

	if userB.Id == 0 {
		if len(userB.PublicId) == 36 {
			_, err = r.db.Exec("insert into social.user_friendship_link (user_id_a, user_id_b, comment, created_at)"+
				"value (?, (select u.id from social.user u where u.public_id = ? limit 1), ?, now())", userA.Id, userB.PublicId, comment)
		} else {
			return fmt.Errorf("UserRepository.CreateFriendshipLink: there no one Id for userB")
		}

	} else {
		_, err = r.db.Exec("insert into social.user_friendship_link (user_id_a, user_id_b, comment, created_at)"+
			"value (?, ?, ?, now())", userA.Id, userB.Id, comment)
	}
	if err != nil {
		if strings.Contains(err.Error(), "1062") {
			return fmt.Errorf("UserRepository.CreateFriendshipLink %d -> %d: already exists", userA.Id, userB.Id)
		}
		return fmt.Errorf("UserRepository.CreateFriendshipLink %d -> %d: %v", userA.Id, userB.Id, err)
	}
	return nil
}

// ForceUserDelete Функция для полного удаления пользователя. В первую очередь для тестов,
// но возможно нужно будет использовать для исполнения закона о цифровом забвении № 149-ФЗ
func (r *UserRepository) ForceUserDelete(userModel *model.UserModel) error {
	interestRepo := GetInterestRepositoryDB(r.db)
	err := interestRepo.UnlinkAllUserInterests(userModel.Id)
	if err != nil {
		return err
	}
	_, err = r.db.Exec("delete from social.user where id = ?", userModel.Id)
	if err != nil {
		return err
	}
	_, err = r.db.Exec("delete from social.user_friendship_link where user_id_a = ? OR user_id_b = ?", userModel.Id, userModel.Id)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetByEmail(email string) (user *model.UserModel, err error) {
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
			"where usr.email = ?", &email)
	userModel := model.UserModel{}
	if err := row.Scan(&userModel.Id, &userModel.PublicId, &userModel.PasswordHash, &userModel.Email,
		&userModel.FirstName, &userModel.LastName, &userModel.MiddleName, &userModel.Gender, &userModel.GenderDesc, &userModel.Town,
		&userModel.CreatedAt, &userModel.UpdatedAt, &userModel.DeletedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("UserRepository.GetByEmail %s : no such user", email)
		}
		return nil, fmt.Errorf("UserRepository.GetByEmail %s: %v", email, err)
	}
	return &userModel, nil
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
