package repository

import (
	"database/sql"
	"fmt"
	"github.com/gGerret/otus-social-prj/repository/model"
	"strings"
	"time"
)

type InterestRepository struct {
	BaseRepository
}

func GetInterestsRepository() *InterestRepository {

	if database == nil {
		panic("Error! Database connection is not initialized")
	}
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

func (r *InterestRepository) GetByInterestName(str string) (*model.InterestRawModel, error) {
	var interest model.InterestRawModel
	row := r.db.QueryRow("select id, interest, created_at FROM social.interests WHERE interest = ?", str)
	if err := row.Scan(&interest.Id, &interest.Interest, &interest.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return &interest, fmt.Errorf("InterestRepository.GetByInterestName %s: no such interest", str)
		}
		return &interest, fmt.Errorf("InterestRepository.GetByInterestName %d: %v", str, err)
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

func (r *InterestRepository) GetInterestsFromList(interests []string) (known []model.InterestRawModel, unknown []model.InterestRawModel, err error) {
	if len(interests) == 0 {
		return nil, nil, nil
	}
	args := make([]interface{}, len(interests))
	for i, id := range interests {
		args[i] = id
	}

	stmt := "select id, interest, created_at from social.interests where interest in (?" + strings.Repeat(",?", len(interests)-1) + ")"
	rows, err := r.db.Query(stmt, args...)
	if err != nil {
		return nil, nil, fmt.Errorf("InterestRepository.GetKnownInterestsFromList: %v", err)
	}
	defer rows.Close()

	knownMap := make(map[string]bool, 0)

	for rows.Next() {
		var knownInterest model.InterestRawModel
		if err := rows.Scan(&knownInterest.Id, &knownInterest.Interest, &knownInterest.CreatedAt); err != nil {
			return nil, nil, fmt.Errorf("InterestRepository.GetKnownInterestsFromList: %v", err)
		}
		known = append(known, knownInterest)
		knownMap[knownInterest.Interest] = true
	}
	for _, i := range interests {
		_, k := knownMap[i]
		if !k {
			unknown = append(unknown, model.InterestRawModel{Interest: i, CreatedAt: sql.NullTime{Time: time.Now(), Valid: true}})
		}
	}

	return known, unknown, nil
}

func (r *InterestRepository) InsertInterestsBulk(interests []model.InterestRawModel) error {
	query := "insert into social.interests (interest, created_at) values "

	for i, interest := range interests {
		query += fmt.Sprintf("( %s, %s )", interest.Interest, interest.CreatedAt.Time.Format("2001-12-31 23:59:59"))
		if i < len(interests)-1 {
			query += ","
		}
	}
	_, err := r.db.Exec(query)

	return err
}

func (r *InterestRepository) InsertInterestsSkipExisting(interests []model.InterestRawModel) (newInterests []model.InterestRawModel, err error) {
	for _, interest := range interests {
		newInterest := model.InterestRawModel{Id: 0, Interest: interest.Interest, CreatedAt: sql.NullTime{Time: time.Now(), Valid: true}}
		ri, err := r.InsertInterest(interest.Interest)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate") {
				eInterest, err := r.GetByInterestName(interest.Interest)
				if err != nil {
					return nil, err
				}
				newInterest.Id = eInterest.Id
				newInterest.Interest = eInterest.Interest
				newInterest.CreatedAt = eInterest.CreatedAt
			} else {
				return nil, err
			}
		} else {
			newInterest.Id = ri
		}
		newInterests = append(newInterests, newInterest)
	}
	return newInterests, nil

}

func (r *InterestRepository) InsertInterest(interest string) (id int64, err error) {
	result, err := r.db.Exec("insert into social.interests (interest, created_at) value (?, ?) ", interest, time.Now())
	if err != nil {
		return 0, fmt.Errorf("InterestRepository.SaveNewInterest: %v", err)
	}
	id, err = result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("InterestRepository.SaveNewInterest: %v", err)
	}
	return id, nil
}

func (r *InterestRepository) LinkInterestsToUser(userId int64, interests []model.InterestRawModel) error {
	for _, interest := range interests {
		_, err := r.db.Exec("insert into social.user_interests_link (user_id, interest_id, created_at) value (?, ?, now()) ON DUPLICATE KEY UPDATE user_id=user_id", userId, interest.Id)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *InterestRepository) UnlinkAllUserInterests(userId int64) error {
	_, err := r.db.Exec("delete from social.user_interests_link where user_id = ?", userId)
	return err
}
