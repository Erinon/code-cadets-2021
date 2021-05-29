package bootstrap

import (
	"code-cadets-2021/homework_4/bets_api/cmd/config"
	"code-cadets-2021/homework_4/bets_api/internal/api"
	"code-cadets-2021/homework_4/bets_api/internal/api/controllers"
	dtomappers "code-cadets-2021/homework_4/bets_api/internal/api/controllers/mappers"
	"code-cadets-2021/homework_4/bets_api/internal/api/controllers/validators"
	dbmappers "code-cadets-2021/homework_4/bets_api/internal/domain/mappers"
	"code-cadets-2021/homework_4/bets_api/internal/domain/services"
	"code-cadets-2021/homework_4/bets_api/internal/infrastructure/sqlite"
)

func newBetStatusValidator() *validators.BetStatusValidator {
	return validators.NewBetStatusValidator()
}

func newDbBetMapper() *dbmappers.BetMapper {
	return dbmappers.NewBetMapper()
}

func newBetRepository(dbExecutor sqlite.DatabaseExecutor, betMapper sqlite.BetMapper) *sqlite.BetRepository {
	return sqlite.NewBetRepository(dbExecutor, betMapper)
}

func newBetService(repository services.BetRepository) *services.BetService {
	return services.NewBetService(repository)
}

func newBetDtoMapper() *dtomappers.BetMapper {
	return dtomappers.NewBetMapper()
}

func newController(betStatusValidator controllers.BetStatusValidator, betService controllers.BetService, betMapper controllers.BetMapper) *controllers.Controller {
	return controllers.NewController(betStatusValidator, betService, betMapper)
}

// Api bootstraps the http server.
func Api(dbExecutor sqlite.DatabaseExecutor) *api.WebServer {
	dbBetMapper := newDbBetMapper()
	betRepository := newBetRepository(dbExecutor, dbBetMapper)

	betService := newBetService(betRepository)

	betStatusValidator := newBetStatusValidator()

	betDtoMapper := newBetDtoMapper()
	controller := newController(betStatusValidator, betService, betDtoMapper)

	return api.NewServer(config.Cfg.Api.Port, config.Cfg.Api.ReadWriteTimeoutMs, controller)
}
