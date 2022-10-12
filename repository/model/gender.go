package model

type GenderDictRawModel struct {
	Id        int
	Code      string
	ShortDesc string
	FullDesc  string
}

func (g *GenderDictRawModel) TableName() string {
	return "gender"
}
