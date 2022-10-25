package entity

import "github.com/gGerret/otus-social-prj/repository/model"

type GenderEntity struct {
	Id        int    `json:"gender_id"`
	Code      string `json:"code"`
	ShortDesc string `json:"short_desc"`
	FullDesc  string `json:"full_desc"`
}

type GendersArrayEntity struct {
	Genders []GenderEntity `json:"genders"`
}

func (g *GenderEntity) FromModel(m *model.GenderDictRawModel) {
	g.Id = m.Id
	g.Code = m.Code
	g.ShortDesc = m.ShortDesc
	g.FullDesc = m.FullDesc
}

func (g *GendersArrayEntity) FromModelArray(ga []model.GenderDictRawModel) {
	for _, genderModel := range ga {
		genderEntity := GenderEntity{}
		genderEntity.FromModel(&genderModel)
		g.Genders = append(g.Genders, genderEntity)
	}
}
