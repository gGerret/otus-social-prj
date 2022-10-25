package repository

import (
	"database/sql"
	"fmt"
	"github.com/gGerret/otus-social-prj/repository/model"
)

type GenderRepository struct {
	BaseRepository
}

func GetGenderRepository() *GenderRepository {

	if database == nil {
		panic("Error! Database connection is not initialized")
	}
	return GetGenderRepositoryDB(database)
}

func GetGenderRepositoryDB(db *sql.DB) *GenderRepository {
	rep := &GenderRepository{}
	rep.db = db
	return rep
}

func (r *GenderRepository) GetById(genderId int64) (*model.GenderDictRawModel, error) {
	var gender model.GenderDictRawModel
	row := r.db.QueryRow("select id, code, short_desc, full_desc FROM social.gender WHERE id = ?", genderId)
	if err := row.Scan(&gender.Id, &gender.Code, &gender.ShortDesc, &gender.FullDesc); err != nil {
		if err == sql.ErrNoRows {
			return &gender, fmt.Errorf("GenderRepository.GetById %d: unknown gender", genderId)
		}
		return &gender, fmt.Errorf("GenderRepository.GetById %d: %v", genderId, err)
	}
	return &gender, nil
}

func (r *GenderRepository) GetAll() ([]model.GenderDictRawModel, error) {

	rows, err := r.db.Query("select id, code, short_desc, full_desc from social.gender order by id limit 3")
	if err != nil {
		return nil, fmt.Errorf("GenderRepository.GetAll: %v", err)
	}
	defer rows.Close()
	var genders []model.GenderDictRawModel

	for rows.Next() {
		var gender model.GenderDictRawModel
		if err := rows.Scan(&gender.Id, &gender.Code, &gender.ShortDesc, &gender.FullDesc); err != nil {
			return nil, fmt.Errorf("GenderRepository.GetAll: %v", err)
		}
		genders = append(genders, gender)
	}
	return genders, nil
}
