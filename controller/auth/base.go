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
	return "base-authenticator"
}
func (b *BaseAuthenticator) GetSystemId() uint {
	return 1
}
func (b *BaseAuthenticator) Do(code string, redirectUri string) (*entity.UserEntity, error) {
	return entity.CreateUserEntityMoc(), nil
}
func (b *BaseAuthenticator) GetAuthURL(string) string {
	return ""
}
