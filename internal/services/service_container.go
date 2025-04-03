package services

import (
	"github.com/Amir-Zouerami/TAPA/internal/repository"
	"github.com/jmoiron/sqlx"
)

type Services struct {
	Dashboard   *DashboardService
	Request     *RequestService
	AppState    *AppStateService
	Environment *EnvironmentService
}

func NewServiceContainer(db *sqlx.DB) *Services {
	metaRepo := repository.NewMetaRepository(db)
	collectionsRepo := repository.NewCollectionsRepository(db)
	environmentsRepo := repository.NewEnvironmentsRepository(db)

	return &Services{
		Dashboard:   NewDashboardService(collectionsRepo, metaRepo, environmentsRepo),
		Request:     NewRequestService(collectionsRepo, metaRepo, environmentsRepo),
		AppState:    NewAppStateService(collectionsRepo, metaRepo, environmentsRepo),
		Environment: NewEnvironmentService(collectionsRepo, metaRepo, environmentsRepo),
	}
}
