package services

import (
	"sync"

	"github.com/Amir-Zouerami/TAPA/internal/models"
	"github.com/Amir-Zouerami/TAPA/internal/repository"
	"github.com/jmoiron/sqlx"
)

type DashboardResult struct {
	Result models.FullDashboardRequestList
	Err    error
}

type DashboardService struct {
	repo *repository.CollectionsRepository
}

// getFullRequestList runs three parallel DB calls, then merges collections, folders, and requests.
// It also extracts requests that do not belong to any folder or collection.
func (s *DashboardService) GetFullRequestList() (models.FullDashboardRequestList, error) {
	var (
		collections    []models.Collection
		folders        []models.Folder
		requests       []models.RequestBasic
		errCollections error
		errFolders     error
		errRequests    error
	)

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		collections, errCollections = s.repo.GetAllCollections()
	}()

	go func() {
		defer wg.Done()
		folders, errFolders = s.repo.GetAllFolders()
	}()

	go func() {
		defer wg.Done()
		requests, errRequests = s.repo.GetAllRequestSummaries()
	}()

	wg.Wait()

	if errCollections != nil {
		return models.FullDashboardRequestList{}, errCollections
	}
	if errFolders != nil {
		return models.FullDashboardRequestList{}, errFolders
	}
	if errRequests != nil {
		return models.FullDashboardRequestList{}, errRequests
	}

	collectionsMap := make(map[int]*models.PopulatedCollection)
	for _, col := range collections {
		collectionsMap[col.ID] = &models.PopulatedCollection{
			Collection: col,
			Folders:    []models.PopulatedFolder{},
			Requests:   []models.RequestBasic{},
		}
	}

	foldersMap := make(map[int]*models.PopulatedFolder)
	for _, folder := range folders {
		ff := models.PopulatedFolder{
			Folder:   folder,
			Requests: []models.RequestBasic{},
		}
		foldersMap[folder.ID] = &ff

		if col, ok := collectionsMap[folder.CollectionID]; ok {
			col.Folders = append(col.Folders, ff)
		}
	}

	looseRequests := []models.RequestBasic{}

	for _, req := range requests {
		if req.FolderID != nil {
			if folder, ok := foldersMap[*req.FolderID]; ok {
				folder.Requests = append(folder.Requests, req)
			}
		} else if req.CollectionID != nil {
			if col, ok := collectionsMap[*req.CollectionID]; ok {
				col.Requests = append(col.Requests, req)
			}
		} else {
			looseRequests = append(looseRequests, req)
		}
	}

	for _, col := range collectionsMap {
		updatedFolders := make([]models.PopulatedFolder, 0, len(col.Folders))
		for _, folder := range col.Folders {
			if updated, ok := foldersMap[folder.Folder.ID]; ok {
				updatedFolders = append(updatedFolders, *updated)
			}
		}
		col.Folders = updatedFolders
	}

	collectionsSlice := make([]models.PopulatedCollection, 0, len(collectionsMap))
	for _, col := range collectionsMap {
		collectionsSlice = append(collectionsSlice, *col)
	}

	return models.FullDashboardRequestList{
		Collections:   collectionsSlice,
		LooseRequests: looseRequests,
	}, nil
}

func NewDashboardService(db *sqlx.DB) *DashboardService {
	return &DashboardService{
		repo: repository.NewCollectionsRepository(db),
	}
}
