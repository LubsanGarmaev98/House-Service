package cmd

import (
	"github.com/timurzdev/mentorship-test-task/internal/handler/create_house"
	"github.com/timurzdev/mentorship-test-task/internal/handler/dummy_login"
	"github.com/timurzdev/mentorship-test-task/internal/handler/middlewares/prometheus"
	"github.com/timurzdev/mentorship-test-task/internal/handler/server"
	"github.com/timurzdev/mentorship-test-task/internal/repository"
	"github.com/timurzdev/mentorship-test-task/internal/service/roles"
	"github.com/timurzdev/mentorship-test-task/internal/service/token"
	create_house_usecase "github.com/timurzdev/mentorship-test-task/internal/usecase/create_house"
)

// контейнер внутренних зависимостей
type Internal struct {
	//external
	*Container

	repository     *repository.Repository
	testRepository *repository.Repository

	server *server.Server

	tokenService  *token.TokenService
	rolesProvider *roles.RoleProvider

	//handlers
	createHouseHandler *create_house.Handler

	dummyLogin *dummy_login.Handler

	//usecases
	createHouseUsecase *create_house_usecase.Usecase

	//middlewares
	prometheusMiddleware *prometheus.Middleware
}

func NewInternal(container *Container) *Internal {
	return &Internal{Container: container}
}

func (i *Internal) GetRepository() *repository.Repository {
	if i.repository == nil {
		i.repository = repository.NewRepository(i.GetPostgres())
	}

	return i.repository
}

func (i *Internal) GetServer() *server.Server {
	if i.server == nil {
		i.server = server.NewServer(
			i.GetLogger(),
			i.configuration.GetServerConfiguration().GetAddress(),
			i.GetCreateHouseHandler(),
			i.GetDummyLoginHandler(),
			i.GetPrometheusMiddleware(),
		)
	}

	return i.server
}

func (i *Internal) GetCreateHouseHandler() *create_house.Handler {
	if i.createHouseHandler == nil {
		i.createHouseHandler = create_house.NewHandler(
			i.GetCreateHouseUsecase(),
			i.GetLogger(),
		)
	}

	return i.createHouseHandler
}

func (i *Internal) GetCreateHouseUsecase() *create_house_usecase.Usecase {
	if i.createHouseUsecase == nil {
		i.createHouseUsecase = create_house_usecase.NewUsecase(i.GetRepository())
	}

	return i.createHouseUsecase
}

func (i *Internal) GetPrometheusMiddleware() *prometheus.Middleware {
	if i.prometheusMiddleware == nil {
		i.prometheusMiddleware = prometheus.New(i.GetMetrics())
	}

	return i.prometheusMiddleware
}

func (i *Internal) GetRolesProvider() *roles.RoleProvider {
	if i.rolesProvider == nil {
		i.rolesProvider = roles.NewRoleProvider()
	}

	return i.rolesProvider
}

func (i *Internal) GetDummyLoginHandler() *dummy_login.Handler {
	if i.dummyLogin == nil {
		i.dummyLogin = dummy_login.NewHandler(
			i.GetTokenService(),
		)
	}

	return i.dummyLogin
}

func (i *Internal) GetTokenService() *token.TokenService {
	if i.tokenService == nil {
		tokenConfigurations := i.configuration.GetTokenConfiguration()
		mySigningKey := []byte(tokenConfigurations.secretKey)
		i.tokenService = token.NewTokenService(mySigningKey, tokenConfigurations.tokenTTL)
	}

	return i.tokenService
}
