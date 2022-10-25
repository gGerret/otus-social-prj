package model

type GenderDictRawModel struct {
	Id        int
	Code      string
	ShortDesc string
	FullDesc  string
}

type GenderModel struct {
	Code      string
	ShortDesc string
	FullDesc  string
}

func (g *GenderDictRawModel) TableName() string {
	return "gender"
}
