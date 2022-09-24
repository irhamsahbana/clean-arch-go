package route

import (
	"ca-boilerplate/bootstrap"
	custom "ca-boilerplate/lib/custom_type"
	"fmt"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_userRoleHttp "ca-boilerplate/domain/user_role/delivery/http"
	_userRoleRepo "ca-boilerplate/domain/user_role/repository/mongo"
	_userRoleUsecase "ca-boilerplate/domain/user_role/usecase"

	_userHttp "ca-boilerplate/domain/user/delivery/http"
	_userRepo "ca-boilerplate/domain/user/repository/mongo"
	_userUsecase "ca-boilerplate/domain/user/usecase"

	_bookHttp "ca-boilerplate/domain/book/delivery/http"
	_bookRepo "ca-boilerplate/domain/book/repository/slice"
	_bookUsecase "ca-boilerplate/domain/book/usecase"

	_fileHttp "ca-boilerplate/domain/file/delivery/http"
	_fileRepo "ca-boilerplate/domain/file/repository/mongo"
	_fileUsecase "ca-boilerplate/domain/file/usecase"

	_tokenRepo "ca-boilerplate/domain/token/repository/mongo"
)

func NewHttpRoutes(r *gin.Engine) {
	if !bootstrap.App.Config.GetBool("app.debug") {
		gin.SetMode(gin.ReleaseMode)
	}

	timeoutContext := time.Duration(bootstrap.App.Config.GetInt("context.timeout")) * time.Second
	mongoDatabase := bootstrap.App.Mongo.Database(bootstrap.App.Config.GetString("mongo.name"))
	appStorageURL := bootstrap.App.Config.GetString("app.url") + bootstrap.App.Config.GetString("app.static_asssets_url")

	r.Static(bootstrap.App.Config.GetString("app.static_asssets_url"), bootstrap.App.Config.GetString("app.static_assets"))
	r.Use(cors.Default())

	tokenRepo := _tokenRepo.NewTokenMongoRepository(*mongoDatabase, custom.TokenableType_USER)
	userRepo := _userRepo.NewUserMongoRepository(*mongoDatabase)
	userRoleRepo := _userRoleRepo.NewUserRoleMongoRepository(*mongoDatabase)
	userUsecase := _userUsecase.NewUserUsecase(userRepo, userRoleRepo, tokenRepo, timeoutContext)
	_userHttp.NewUserHandler(r, userUsecase)

	userRoleUsecase := _userRoleUsecase.NewUserRoleUsecase(userRoleRepo, timeoutContext)
	_userRoleHttp.NewUserRoleHandler(r, userRoleUsecase)

	fileRepo := _fileRepo.NewFileMongoRepository(*mongoDatabase)
	fileUsecase := _fileUsecase.NewFileUploadUsecase(fileRepo, timeoutContext)
	_fileHttp.NewFileHandler(r, fileUsecase, appStorageURL)

	bookRepo := _bookRepo.NewBookSliceRepository()
	bookUsecase := _bookUsecase.NewBookUsecase(bookRepo)
	_bookHttp.NewBookHandler(r, bookUsecase)

	appPort := fmt.Sprintf(":%v", bootstrap.App.Config.GetString("server.address"))
	log.Fatal(r.Run(appPort))
}
