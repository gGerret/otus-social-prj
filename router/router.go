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

	serverApiPath := q.cfg.Api.ApiURL + q.cfg.Api.Version

	//Add auth filter
	q.guard = jwt.NewGuard(
		q.cfg.Auth.Guard,
		logger.Named("guard"),
		//Эндпоинты, для которых не проверяется наличие токена аутентификации
		serverApiPath+"/register",
		serverApiPath+q.cfg.Auth.AuthUrl+"/login",
		serverApiPath+q.cfg.Auth.AuthUrl+"/test_init", //TODO: Надо будет убрать для прода
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

	apiRoute := q.router.Group(serverApiPath)
	{
		//Всё про авторизацию. Выделяем в отдельную группу.
		authRoute := apiRoute.Group(q.cfg.Auth.AuthUrl)
		{

			authRoute.POST("/login", q.authCtrl.PostUserPass)
			authRoute.GET("/test_init", q.testCtrl.InitTestDB)
		}
		//Все справочники
		dictionaryRoute := apiRoute.Group("/dict")
		{
			dictionaryRoute.GET("/gender/all", q.dictCtrl.GetKnownGenders)
			dictionaryRoute.GET("/interest/all", q.dictCtrl.GetKnownInterests)
		}
		//Регистрация - отдельная вселенная
		apiRoute.POST("/register", q.userCtrl.RegisterUser)
		apiRoute.DELETE("/registered", q.userCtrl.DeleteCurrentUser)

		//Про пользователя
		userRoute := apiRoute.Group("/user")
		{
			userRoute.GET("/", q.userCtrl.GetCurrentUser)
			userRoute.GET("/page", q.userCtrl.GetCurrentUserPage)
			userRoute.GET("/friendship", q.userCtrl.GetCurrentUserFriends)
			userRoute.GET("/:id", q.userCtrl.GetUserById)
			userRoute.GET("/:id/page", q.userCtrl.GetUserPage)

			userRoute.PUT("/", q.userCtrl.UpdateCurrentUser)
			userRoute.PUT("/page", q.userCtrl.UpdateCurrentUserPage)
			userRoute.PUT("/friendship", q.userCtrl.MakeFriendship)

			userRoute.POST("/query", q.userCtrl.GetUserByFilter)
		}

	}

	return q, nil
}

func (q *SocialServer) GetBaseURI() string {
	return fmt.Sprintf("%s:%d", q.cfg.BaseURL, q.cfg.Port)
}

func (q *SocialServer) RunServer() error {
	q.logger.Info("Starting server in " + q.cfg.Mode + " mode...")
	gin.SetMode(q.cfg.Mode)
	defer func() {
		q.logger.Info("Server shutdown")
	}()
	return q.router.Run(fmt.Sprintf(":%d", q.cfg.Port))
}
