package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/opensourceways/server-common-lib/interrupts"
	"github.com/sirupsen/logrus"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/opensourceways/app-cla-signing/cla/app"
	"github.com/opensourceways/app-cla-signing/cla/controller"
	"github.com/opensourceways/app-cla-signing/cla/domain/corpemaildomainemail"
	"github.com/opensourceways/app-cla-signing/cla/domain/emailclient"
	"github.com/opensourceways/app-cla-signing/cla/domain/signingcodeemail"
	"github.com/opensourceways/app-cla-signing/cla/domain/userservice"
	"github.com/opensourceways/app-cla-signing/cla/domain/vcservice"
	"github.com/opensourceways/app-cla-signing/cla/infrastructure/corpemaildomainemailimpl"
	"github.com/opensourceways/app-cla-signing/cla/infrastructure/emailclientimpl"
	"github.com/opensourceways/app-cla-signing/cla/infrastructure/emaildeliveryimpl"
	"github.com/opensourceways/app-cla-signing/cla/infrastructure/encryptionimpl"
	"github.com/opensourceways/app-cla-signing/cla/infrastructure/passwordimpl"
	"github.com/opensourceways/app-cla-signing/cla/infrastructure/randomcodeimpl"
	"github.com/opensourceways/app-cla-signing/cla/infrastructure/repositoryimpl"
	"github.com/opensourceways/app-cla-signing/cla/infrastructure/signingcodeemailimpl"
	"github.com/opensourceways/app-cla-signing/cla/infrastructure/tokenimpl"
	"github.com/opensourceways/app-cla-signing/common/infrastructure/mongodb"
	"github.com/opensourceways/app-cla-signing/server/config"
	"github.com/opensourceways/app-cla-signing/server/docs"
)

func StartWebServer(port int, timeout time.Duration, cfg *config.Config) {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(logRequest())

	setRouter(r, cfg)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: r,
	}

	defer interrupts.WaitForGracefulShutdown()

	interrupts.ListenAndServe(srv, timeout)
}

func logRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		endTime := time.Now()

		logrus.Infof(
			"| %d | %d | %s | %s |",
			c.Writer.Status(),
			endTime.Sub(startTime),
			c.Request.Method,
			c.Request.RequestURI,
		)
	}
}

//setRouter init router
func setRouter(engine *gin.Engine, cfg *config.Config) {
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Title = "CLA"
	docs.SwaggerInfo.Description = "set header: 'PRIVATE-TOKEN=xxx'"

	v1 := engine.Group(docs.SwaggerInfo.BasePath)
	setApiV1(v1, cfg)

	engine.UseRawPath = true
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

func setApiV1(v1 *gin.RouterGroup, cfg *config.Config) {
	randomCode := randomcodeimpl.NewRandomCodeImpl()

	repo := repositoryimpl.NewVerificationCode(
		mongodb.DAO(cfg.Mongodb.Collections.VerificationCode),
	)

	vcService := vcservice.NewVCService(repo, randomCode)

	initVerificationCode(v1, cfg, vcService)
	initEmployeeManager(v1, cfg, vcService)
	initEmployeeSigning(v1, cfg)
}

func initVerificationCode(v1 *gin.RouterGroup, cfg *config.Config, vcService vcservice.VCService) {
	cli := emailclient.NewEmailClient(emailclientimpl.NewEmailClientImpl())

	delivery := emaildeliveryimpl.NewEmailDeliveryImpl()

	signingCode := signingcodeemail.NewSigningCodeEmail(
		signingcodeemailimpl.NewSigningCodeEmailImpl(),
	)

	corpEmailDomain := corpemaildomainemail.NewCorpEmailDomainEmail(
		corpemaildomainemailimpl.NewCorpEmailDomainEmailImpl(),
	)

	// VerificationCodeController
	controller.AddRouteForVerificationCodeController(
		v1, app.NewSigningCodeService(
			cli, signingCode, delivery, vcService,
		),
		app.NewEmailDomainCodeService(
			cli, corpEmailDomain, delivery, vcService,
		),
	)
}

func initEmployeeManager(v1 *gin.RouterGroup, cfg *config.Config, vcService vcservice.VCService) {
	repo := repositoryimpl.NewCorpSigning(
		mongodb.DAO(cfg.Mongodb.Collections.CorpSigning),
	)

	userRepo := repositoryimpl.NewUser(
		mongodb.DAO(cfg.Mongodb.Collections.User),
	)

	encrypt := encryptionimpl.NewEncryptionImpl()

	password := passwordimpl.NewPasswordImpl()

	user := userservice.NewUserService(userRepo, encrypt, password)

	// EmployeeManagerController
	controller.AddRouteForEmployeeManagerController(
		v1, app.NewEmployeeManagerService(
			repo, user,
		),
	)

	// CorpEmailDomainController
	controller.AddRouteForCorpEmailDomainController(
		v1, app.NewCorpEmailDomainService(
			repo, vcService,
		),
	)

	// UserController
	controller.AddRouteForUserController(
		v1, app.NewUserService(repo, user),
		tokenimpl.NewTokenImpl(),
	)
}

func initEmployeeSigning(v1 *gin.RouterGroup, cfg *config.Config) {
	repo := repositoryimpl.NewEmployeeSigning(
		mongodb.DAO(cfg.Mongodb.Collections.EmployeeSigning),
	)

	controller.AddRouteForEmployeeSigningController(
		v1, app.NewEmployeeSigningService(repo),
	)
}
