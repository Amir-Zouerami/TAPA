package services

import (
	"github.com/Amir-Zouerami/TAPA/internal/repository"
)

type AppStateService struct {
	collectionsRepo  repository.CollectionsRepositoryInterface
	environmentsRepo repository.EnvironmentsRepositoryInterface
	metaRepo         repository.MetaRepositoryInterface
}

func NewAppStateService(
	collectionsRepo repository.CollectionsRepositoryInterface,
	metaRepo repository.MetaRepositoryInterface,
	environmentsRepo repository.EnvironmentsRepositoryInterface,
) *AppStateService {
	return &AppStateService{
		collectionsRepo:  collectionsRepo,
		metaRepo:         metaRepo,
		environmentsRepo: environmentsRepo,
	}
}
