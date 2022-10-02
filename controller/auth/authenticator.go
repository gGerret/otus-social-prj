package auth

import (
	"github.com/gGerret/otus-social-prj/controller/entity"
	"go.uber.org/zap"
)

type Authenticator interface {
	GetName() string
	GetSystemId() uint
	Do(code string, redirectUri string) (*entity.UserEntity, error)
	GetAuthURL(string) string
}

func GetAuthenticator(authType string, logger *zap.SugaredLogger) (Authenticator, *entity.ErrorEntity) {
	var a Authenticator
	switch authType {
	case "base":
		a = NewBaseAuthenticator(logger)
	case "vk":
		return nil, &entity.ErrNotImplemented
		break
	case "google":
		return nil, &entity.ErrNotImplemented
	default:
		return nil, &entity.AuthErrWrongType
	}
	return a, nil
}
