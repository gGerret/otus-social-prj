package jwt

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gGerret/otus-social-prj/repository"
	"github.com/gGerret/otus-social-prj/social"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"net/http"
	"time"
)

const DefaultHeader = "X-Auth-Token"

type Guard struct {
	cfg    *ConfigGuard
	logger *social.SocialLogger
	User   Authenticatable
	exSet  map[string]struct{}
}

type Authenticatable interface {
	LoadViaJwtSub(sub interface{})
	GetJwtClaims() *jwt.MapClaims
}

func NewGuard(cfg *ConfigGuard, logger *social.SocialLogger, exceptionList ...string) *Guard {
	guard := &Guard{
		cfg:    cfg,
		logger: logger,
	}
	guard.exSet = make(map[string]struct{}, len(exceptionList))

	for _, s := range exceptionList {
		guard.exSet[s] = struct{}{}
	}

	logger.Debug(guard.exSet)
	return guard
}
func (g *Guard) isException(uri string) bool {
	_, ok := g.exSet[uri]
	return ok
}
func (g *Guard) AuthFilter(ctx *gin.Context) {
	localLogger := g.logger.ContextLogger(ctx.GetString("reqId"), "AuthFilter")
	header := g.cfg.Header
	if header == "" {
		header = DefaultHeader
	}

	if ctx.Request.Method == "OPTIONS" {
		ctx.Next()
		return
	}
	xAuthToken := ctx.Request.Header.Get(header)
	localLogger.Debug(ctx.Request.URL.Path)

	if g.isException(ctx.Request.URL.Path) {
		ctx.Next()
		return
	}

	localLogger.Debugf("xAuthToken = %s", xAuthToken)
	if xAuthToken == "" {
		g.abort(ctx, errors.New("token is empty"))
		return
	}
	token, err := jwt.ParseWithClaims(xAuthToken, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		secret := g.cfg.Secret
		if secret == "" {
			panic("Secret for JWT auth is not defined!")
		}

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secret), nil
	})

	if err != nil {
		g.logger.Error("token parse error ", err)
		g.abort(ctx, err)
		return
	}
	g.logger.Debug("token ", token)

	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		if claims.Subject != "" {
			publicUid, err := uuid.FromString(claims.Subject)
			if err != nil {
				g.logger.Error("get uuid error ", err)
				g.abort(ctx, err)
				return
			}
			rep := repository.GetUserRepository()
			user, err := rep.GetByPublicId(publicUid.String())
			if err != nil {
				g.logger.Error("get user error ", err)
				g.abort(ctx, err)
				return
			}

			ctx.Set("User", user)

		} else {
			g.logger.Error("claims.Subject == \"\"")
			g.abort(ctx, errors.New("subject in token is empty"))
			return
		}

	} else {
		g.logger.Error("claims is not valid ", claims, err)

		g.abort(ctx, errors.New("claims is not valid"))
		return
	}

	ctx.Next()
}

func (g *Guard) SetToken(publicId string, authType string) (string, error) {
	// Create the Claims
	claims := &jwt.StandardClaims{
		Subject:   publicId,
		ExpiresAt: time.Now().UTC().Add(time.Duration(g.cfg.TokenLifeHours) * time.Hour).Unix(),
		NotBefore: time.Now().UTC().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(g.cfg.Secret))
	return tokenString, err
}

func (g *Guard) abort(ctx *gin.Context, err ...error) {
	if len(err) == 0 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	} else {
		ctx.AbortWithError(http.StatusUnauthorized, err[0])
	}
}
