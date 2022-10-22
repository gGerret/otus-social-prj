package model

import (
	"database/sql"
)

type InterestRawModel struct {
	Id        int64
	Interest  string
	CreatedAt sql.NullTime
}

type UserInterestLinkRawModel struct {
	UserId     int64
	InterestId int64
}

type UserInterestsModel struct {
	UserId    int64
	Interests []string
}

func (i *InterestRawModel) TableName() string {
	return "interests"
}

func (iul *UserInterestLinkRawModel) TableName() string {
	return "user_interests_link"
}
