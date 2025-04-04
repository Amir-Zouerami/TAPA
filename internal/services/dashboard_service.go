package services

import (
	"context"
	"os"
	"path/filepath"
	"time"

	"github.com/Amir-Zouerami/TAPA/internal/errors"
	"github.com/Amir-Zouerami/TAPA/internal/models"
	"github.com/Amir-Zouerami/TAPA/internal/repository"
	"golang.org/x/sync/errgroup"
)

type DashboardService struct {
	collectionsRepo  repository.CollectionsRepositoryInterface
	environmentsRepo repository.EnvironmentsRepositoryInterface
	metaRepo         repository.MetaRepositoryInterface
}

// GetCollectionsTree runs three parallel DB calls to return collections, folders and the request maps
func (s *DashboardService) GetCollectionsTree() (models.CollectionsTree, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	group, ctx := errgroup.WithContext(ctx)
	colTree := models.CollectionsTree{}

	group.Go(func() error {
		collections, err := s.collectionsRepo.GetAllCollections(ctx)

		if err != nil {
			return err
		}

		collectionsMap := make(map[int]models.Collection, len(collections))
		for _, col := range collections {
			collectionsMap[col.ID] = col
		}

		colTree.Collections = collectionsMap
		return nil
	})

	group.Go(func() error {
		folders, err := s.collectionsRepo.GetAllFolders(ctx)

		if err != nil {
			return err
		}

		foldersMap := make(map[int]models.Folder, len(folders))
		for _, folder := range folders {
			foldersMap[folder.ID] = folder
		}

		colTree.Folders = foldersMap
		return nil
	})

	group.Go(func() error {
		requestList, err := s.collectionsRepo.GetAllRequestSummaries(ctx)

		if err != nil {
			return err
		}

		requestsMap := make(map[int]models.RequestSummary, len(requestList))
		for _, request := range requestList {
			requestsMap[request.ID] = request
		}

		colTree.RequestList = requestsMap
		return nil
	})

	if err := group.Wait(); err != nil {
		return models.CollectionsTree{}, errors.Wrap(errors.ErrLoadingCollectionTree, err)
	}

	return colTree, nil
}

// Decides if the app is launched for the first time on user's computer
func (s *DashboardService) IsFirstLaunch() (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	isFirstLaunch, err := s.metaRepo.IsFirstLaunch(ctx)

	if err != nil {
		return false, errors.Wrap(errors.ErrFirstLaunchRead, err)
	}

	return isFirstLaunch, nil
}

// Set the first launch to done (called once after the first launch)
func (s *DashboardService) SetFirstLaunchDone() (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := s.metaRepo.SetFirstLaunchDone(ctx)

	if err != nil {
		return false, errors.Wrap(errors.ErrUpdatingFirstLaunch, err)
	}

	return true, nil
}

// Fetches the user config directory based on their OS
func (s *DashboardService) GetUserConfigDir() (string, error) {
	configDir, err := os.UserConfigDir()

	if err != nil {
		return "", errors.Wrap(errors.ErrGetUserConfigDirectory, err)
	}

	appDir := filepath.Join(configDir, "tapa")
	return appDir, nil
}

// LoadDashboardData fetches the initial data required to bootstrap the dashboard.
func (s *DashboardService) LoadDashboardData() (models.DashboardData, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	dashboardData := models.DashboardData{}
	group, ctx := errgroup.WithContext(ctx)

	group.Go(func() error {
		environments, err := s.environmentsRepo.ListAllEnvironments(ctx)
		if err != nil {
			return err
		}

		environmentsMap := make(map[int]models.Environment, len(environments))
		for _, env := range environments {
			environmentsMap[env.ID] = env
		}

		dashboardData.Environments = environmentsMap
		return nil
	})

	group.Go(func() error {
		keyboardShortcuts, err := s.metaRepo.GetKeyboardShortcuts(ctx)
		if err != nil {
			return err
		}

		result := make(map[int]models.KeyboardShortcut)
		for _, shortcut := range keyboardShortcuts {
			result[shortcut.ID] = shortcut
		}

		dashboardData.KeyboardShortcuts = result
		return nil
	})

	group.Go(func() error {
		userSettings, err := s.metaRepo.GetUserSettings(ctx)
		if err != nil {
			return err
		}

		dashboardData.UserSettings = userSettings
		return nil
	})

	group.Go(func() error {
		appState, err := s.metaRepo.GetLastAppState(ctx)
		if err != nil {
			return err
		}

		if err = appState.AfterLoad(); err != nil {
			return err
		}

		dashboardData.AppState = appState
		return nil
	})

	if err := group.Wait(); err != nil {
		return models.DashboardData{}, errors.Wrap(errors.ErrLoadingDashboardData, err)
	}

	if dashboardData.AppState.SelectedEnvironment != nil {
		environmentMap, err := s.environmentsRepo.GetEnvironmentDetails(ctx, *dashboardData.AppState.SelectedEnvironment)

		if err != nil {
			return models.DashboardData{}, errors.Wrap(errors.ErrLoadingDashboardData, err)
		}

		result := make(map[string]string)
		for _, variable := range environmentMap {
			result[variable.Key] = variable.Value
		}

		dashboardData.CurrentEnvironmentVariables = result
	}

	return dashboardData, nil
}

func NewDashboardService(
	collectionsRepo repository.CollectionsRepositoryInterface,
	metaRepo repository.MetaRepositoryInterface,
	environmentsRepo repository.EnvironmentsRepositoryInterface,
) *DashboardService {
	return &DashboardService{
		collectionsRepo:  collectionsRepo,
		metaRepo:         metaRepo,
		environmentsRepo: environmentsRepo,
	}
}
