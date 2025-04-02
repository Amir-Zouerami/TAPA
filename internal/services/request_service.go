package services

import (
	"github.com/Amir-Zouerami/TAPA/internal/repository"
)

type RequestService struct {
	collectionsRepo  repository.CollectionsRepositoryInterface
	environmentsRepo repository.EnvironmentsRepositoryInterface
	metaRepo         repository.MetaRepositoryInterface
}

func NewRequestService(
	collectionsRepo repository.CollectionsRepositoryInterface,
	metaRepo repository.MetaRepositoryInterface,
	environmentsRepo repository.EnvironmentsRepositoryInterface,
) *RequestService {
	return &RequestService{
		collectionsRepo:  collectionsRepo,
		metaRepo:         metaRepo,
		environmentsRepo: environmentsRepo,
	}
}
