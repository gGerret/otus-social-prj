package controller

import (
	"github.com/gGerret/otus-social-prj/controller/entity"
	"github.com/gGerret/otus-social-prj/repository"
	"github.com/gGerret/otus-social-prj/social"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DictionaryController struct {
	ApiController
}

func NewDictionaryController(config *ConfigApi, logger *social.SocialLogger) *DictionaryController {
	c := &DictionaryController{}
	c.Init(config, logger)
	c.Name = "DictionaryController"
	return c
}

func (c *DictionaryController) GetKnownInterests(ctx *gin.Context) {
	localLogger := c.logger.ContextLogger(ctx.GetString("reqId"), "GetKnownInterests")
	ec := NewErrHelper(ctx, localLogger)

	rep := repository.GetInterestRepository()

	interests, err := rep.GetAll()
	if err != nil {
		ec.SetErr(entity.ErrNotFound, err)
		return
	} else {
		interestsEntity := &entity.InterestsArrayEntity{}
		interestsEntity.FromStringArray(interests)
		ctx.JSON(http.StatusOK, interestsEntity)
	}
}

func (c *DictionaryController) GetKnownGenders(ctx *gin.Context) {
	localLogger := c.logger.ContextLogger(ctx.GetString("reqId"), "GetKnownGenders")
	ec := NewErrHelper(ctx, localLogger)

	rep := repository.GetGenderRepository()

	genders, err := rep.GetAll()
	if err != nil {
		ec.SetErr(entity.ErrNotFound, err)
		return
	} else {
		gendersEntity := &entity.GendersArrayEntity{}
		gendersEntity.FromModelArray(genders)
		ctx.JSON(http.StatusOK, gendersEntity)
	}
}
