package repository

import (
	"database/sql"
	"fmt"
	"github.com/gGerret/otus-social-prj/repository/model"
	"time"
)

type InterestRepository struct {
	BaseRepository
}

func GetInterestsRepository() *InterestRepository {
	//Пока без БД балуемся
	/*if database == nil {
	    panic("Error! Database connection is not initialized")
	}*/
	return GetInterestRepositoryDB(nil)
}

func GetInterestRepositoryDB(db *sql.DB) *InterestRepository {
	rep := &InterestRepository{}
	rep.db = db
	return rep
}

func (r *InterestRepository) GetById(interestId int64) (*model.InterestRawModel, error) {
	var interest model.InterestRawModel
	row := r.db.QueryRow("select id, interest, created_at FROM social.interests WHERE id = ?", interestId)
	if err := row.Scan(&interest.Id, &interest.Interest, &interest.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return &interest, fmt.Errorf("InterestRepository.GetById %d: no such interest", interestId)
		}
		return &interest, fmt.Errorf("InterestRepository.GetById %d: %v", interestId, err)
	}
	return &interest, nil
}

func (r *InterestRepository) GetByUserId(userId int64) (*model.UserInterestsModel, error) {
	rows, err := r.db.Query(""+
		"select i.interest from social.user_interests_link uil "+
		"   join social.interests i on uil.interest_id = i.id "+
		"where uil.user_id = ?", userId)
	if err != nil {
		return nil, fmt.Errorf("InterestRepository.GetByUserId %d: %v", userId, err)
	}
	defer rows.Close()
	var userInterests []string

	for rows.Next() {
		var interest string
		if err := rows.Scan(&interest); err != nil {
			return nil, fmt.Errorf("InterestRepository.GetByUserId %d: %v", userId, err)
		}
		userInterests = append(userInterests, interest)
	}
	return &model.UserInterestsModel{UserId: userId, Interests: userInterests}, nil
}

func (r *InterestRepository) GetKnownInterestsFromList(interests []string) ([]model.InterestRawModel, error) {
	rows, err := r.db.Query(""+
		"select id, interest, created_at from social.interests where interest in (?)", interests)
	if err != nil {
		return nil, fmt.Errorf("InterestRepository.GetKnownInterestsFromList: %v", err)
	}
	defer rows.Close()
	var knownInterests []model.InterestRawModel

	for rows.Next() {
		var knownInterest model.InterestRawModel
		if err := rows.Scan(&knownInterest.Id, &knownInterest.Interest, &knownInterest.CreatedAt); err != nil {
			return nil, fmt.Errorf("InterestRepository.GetKnownInterestsFromList: %v", err)
		}
		knownInterests = append(knownInterests, knownInterest)
	}
	return knownInterests, nil
}
func (r *InterestRepository) GetKnownStrInterestsFromList(interests []string) ([]string, error) {
	rows, err := r.db.Query(""+
		"select interest from social.interests where interest in (?)", interests)
	if err != nil {
		return nil, fmt.Errorf("InterestRepository.GetKnownStrInterestsFromList: %v", err)
	}
	defer rows.Close()
	var userInterests []string

	for rows.Next() {
		var interest string
		if err := rows.Scan(&interest); err != nil {
			return nil, fmt.Errorf("InterestRepository.GetKnownStrInterestsFromList %v", err)
		}
		userInterests = append(userInterests, interest)
	}
	return userInterests, nil
}

func (r *InterestRepository) SaveNewInterest(interest string) (int64, error) {
	result, err := r.db.Exec("insert into social.interests (interest, created_at) value (?, ?)", interest, time.Now())
	if err != nil {
		return 0, fmt.Errorf("InterestRepository.SaveNewInterest: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("InterestRepository.SaveNewInterest: %v", err)
	}
	return id, nil
}
