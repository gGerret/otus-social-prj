package entity

type InterestsArrayEntity struct {
	Interests []string `json:"interests"`
}

func (i *InterestsArrayEntity) FromStringArray(interests []string) {
	i.Interests = interests
}
