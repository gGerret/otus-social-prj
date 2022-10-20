package model

import "database/sql"

type FriendshipLinkRawModel struct {
	Id         int64
	UserIdA    int64
	UserIdB    int64
	Comment    string
	CreatedAt  sql.NullTime
	ApprovedAt sql.NullTime
}

func (u *FriendshipLinkRawModel) TableName() string {
	return "user_friendship_link"
}
