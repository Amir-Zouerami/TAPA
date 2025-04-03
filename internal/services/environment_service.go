package services

import (
	"context"
	"time"

	"github.com/Amir-Zouerami/TAPA/internal/errors"
	"github.com/Amir-Zouerami/TAPA/internal/repository"
)

type EnvironmentService struct {
	environmentsRepo repository.EnvironmentsRepositoryInterface
	collectionsRepo  repository.CollectionsRepositoryInterface
	metaRepo         repository.MetaRepositoryInterface
}

func (s *EnvironmentService) GetEnvironmentsVariables(envID int) (map[string]string, error) {
	ctx, close := context.WithTimeout(context.Background(), time.Second*5)
	defer close()

	allVariables, err := s.environmentsRepo.GetEnvironmentDetails(ctx, envID)
	if err != nil {
		return nil, errors.Wrap(errors.ErrReadingEnvironmentsVariables, err)
	}

	variablesMap := make(map[string]string)
	for _, variable := range allVariables {
		variablesMap[variable.Key] = variable.Value
	}

	return variablesMap, nil
}

func NewEnvironmentService(
	collectionsRepo repository.CollectionsRepositoryInterface,
	metaRepo repository.MetaRepositoryInterface,
	environmentsRepo repository.EnvironmentsRepositoryInterface,
) *EnvironmentService {
	return &EnvironmentService{
		environmentsRepo: environmentsRepo,
		collectionsRepo:  collectionsRepo,
		metaRepo:         metaRepo,
	}
}
