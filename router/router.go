package router

import (
	"fmt"
	"github.com/gGerret/otus-social-prj/controller"
	"github.com/gGerret/otus-social-prj/controller/auth/jwt"
	"github.com/gGerret/otus-social-prj/social"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type SocialServer struct {
	router *gin.Engine
	logger *social.SocialLogger
	cfg    *ServerConfig

	guard *jwt.Guard

	//Authentication controllers
	authCtrl *controller.AuthController

	//API Controllers
	userCtrl *controller.UserController
	dictCtrl *controller.DictionaryController
	testCtrl *controller.TestController
}

func NewSocialServer(config *ServerConfig, logger *social.SocialLogger) (*SocialServer, error) {

	q := &SocialServer{
		logger: logger,
		router: gin.Default(),
		cfg:    config,
	}
	q.logger.Info("Initializing Social Web server...")

	//Add auth filter
	q.guard = jwt.NewGuard(
		q.cfg.Auth.Guard,
		logger.Named("guard"),
		//Эндпоинты, для которых не проверяется наличие токена аутентификации
		q.cfg.Api.ApiURL+"/register",
		q.cfg.Api.ApiURL+q.cfg.Auth.AuthUrl+"/login",
		q.cfg.Api.ApiURL+q.cfg.Auth.AuthUrl+"/test_init", //TODO: Надо будет убрать для прода
	)

	q.router.Use(controller.BaseFilter)
	q.router.Use(q.guard.AuthFilter)

	//Корсы...
	q.router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"}, //TODO: ОБЯЗАТЕЛЬНО СДЕЛАТЬ, ЧТОБЫ В ПРОДЕ ПОДСТАВЛЯЛСЯ НОРМАЛЬНЫЙ ОРИДЖИН
		AllowMethods: []string{"GET", "POST", "PUT", "OPTIONS"},
		AllowHeaders: []string{"Origin", q.cfg.Auth.Guard.Header, "authorization", "Content-Type"},
		//ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	q.authCtrl = controller.NewAuthController(q.cfg.Auth, logger.Named("auth-controller"), q.guard)

	q.userCtrl = controller.NewUserController(q.cfg.Api, logger.Named("user-controller"))

	q.dictCtrl = controller.NewDictionaryController(q.cfg.Api, logger.Named("dictionary-controller"))

	q.testCtrl = controller.NewTestController(q.cfg.Api, logger.Named("test-controller"))

	apiRoute := q.router.Group(q.cfg.Api.ApiURL)
	{
		//Всё про авторизацию. Выделяем в отдельную группу.
		authRoute := apiRoute.Group(q.cfg.Auth.AuthUrl)
		{

			authRoute.POST("/login", q.authCtrl.PostUserPassMock)
			authRoute.GET("/test_init", q.testCtrl.InitTestDB)
		}
		//Регистрация - отдельная вселенная
		apiRoute.POST("/register", q.userCtrl.RegisterUser)

		//Про пользователя
		apiRoute.GET("/user", q.userCtrl.GetCurrentUser)
		apiRoute.GET("/user/:id", q.userCtrl.GetUserById)
		apiRoute.PUT("/user", q.userCtrl.PutUser)
		apiRoute.POST("user/query", q.userCtrl.GetUserByFilter)
		//apiRoute.GET("/user/friend", q.userCtrl.GetUserFriends)
		//apiRoute.POST("/user/friend/:friend_id", q.userCtrl.MakeFriendship)

		//Справочники
		apiRoute.GET("/dict/interests", q.dictCtrl.GetKnownInterests)
		apiRoute.GET("/dict/genders", q.dictCtrl.GetKnownGenders)
	}

	return q, nil
}

func (q *SocialServer) GetBaseURI() string {
	return fmt.Sprintf("%s:%d", q.cfg.BaseURL, q.cfg.Port)
}

func (q *SocialServer) RunServer() error {
	q.logger.Info("Starting server...")
	defer func() {
		q.logger.Info("Server shutdown")
	}()
	return q.router.Run(fmt.Sprintf(":%d", q.cfg.Port))
}
