package auth

import (
	"github.com/gGerret/otus-social-prj/controller/entity"
	"go.uber.org/zap"
)

type BaseAuthenticator struct {
	logger *zap.SugaredLogger
}

func NewBaseAuthenticator(logger *zap.SugaredLogger) Authenticator {
	a := BaseAuthenticator{logger.Named("base-authenticator")}

	return &a
}
func (b *BaseAuthenticator) GetName() string {

}
func (b *BaseAuthenticator) GetSystemId() uint {

}
func (b *BaseAuthenticator) Do(code string, redirectUri string) (*entity.UserEntity, error) {

}
func (b *BaseAuthenticator) GetAuthURL(string) string {

}
