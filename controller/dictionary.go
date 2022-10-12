package controller

import (
	"github.com/gGerret/otus-social-prj/social"
	"github.com/gin-gonic/gin"
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

}

func (c *DictionaryController) GetKnownGenders(ctx *gin.Context) {

}
