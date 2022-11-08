package entity

type UpdateUserPageEntity struct {
	UserPublicId string `json:"user_id"`
	PagePublicId string `json:"page_id"`
}
