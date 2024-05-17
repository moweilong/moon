// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package palace

import (
	"github.com/aide-cloud/moon/cmd/server/palace/internal/biz"
	"github.com/aide-cloud/moon/cmd/server/palace/internal/data"
	"github.com/aide-cloud/moon/cmd/server/palace/internal/data/repoimpl"
	"github.com/aide-cloud/moon/cmd/server/palace/internal/palaceconf"
	"github.com/aide-cloud/moon/cmd/server/palace/internal/server"
	"github.com/aide-cloud/moon/cmd/server/palace/internal/service"
	"github.com/aide-cloud/moon/cmd/server/palace/internal/service/authorization"
	"github.com/aide-cloud/moon/cmd/server/palace/internal/service/user"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(bootstrap *palaceconf.Bootstrap, logger log.Logger) (*kratos.App, func(), error) {
	grpcServer := server.NewGRPCServer(bootstrap)
	dataData, cleanup, err := data.NewData(bootstrap)
	if err != nil {
		return nil, nil, err
	}
	captchaRepo := repoimpl.NewCaptchaRepo(dataData)
	captchaBiz := biz.NewCaptchaBiz(captchaRepo)
	userRepo := repoimpl.NewUserRepo(dataData)
	teamRepo := repoimpl.NewTeamRepo(dataData)
	cacheRepo := repoimpl.NewCacheRepo(dataData)
	authorizationBiz := biz.NewAuthorizationBiz(userRepo, teamRepo, cacheRepo)
	authorizationService := authorization.NewAuthorizationService(captchaBiz, authorizationBiz)
	httpServer := server.NewHTTPServer(bootstrap, authorizationService)
	greeterRepo := data.NewGreeterRepo(dataData)
	greeterUsecase := biz.NewGreeterUsecase(greeterRepo)
	greeterService := service.NewGreeterService(greeterUsecase)
	transactionRepo := repoimpl.NewTransactionRepo(dataData)
	userBiz := biz.NewUserBiz(userRepo, transactionRepo)
	userService := user.NewUserService(userBiz)
	serverServer := server.RegisterService(grpcServer, httpServer, greeterService, userService, authorizationService)
	app := newApp(bootstrap, serverServer, logger)
	return app, func() {
		cleanup()
	}, nil
}
