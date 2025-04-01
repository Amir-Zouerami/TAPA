package errors

// ------------- Dashboard Service
var (
	ErrLoadingCollectionTree = &TapaError{Code: 4000, Message: "Failed loading collection tree -> "}
	ErrLoadingDashboardData  = &TapaError{Code: 4001, Message: "Failed loading initial dashboard data -> "}
	ErrFetchingLastAppState  = &TapaError{Code: 4002, Message: "Failed fetching the last app state -> "}
	ErrSavingLastAppState    = &TapaError{Code: 4003, Message: "Failed saving the last app state -> "}
	ErrFirstLaunchRead       = &TapaError{Code: 4004, Message: "Failed fetching first launch status -> "}
	ErrUpdatingFirstLaunch   = &TapaError{Code: 4005, Message: "Failed updating the first launch status -> "}
)

// ------------- Environment Service
var (
	ErrReadingEnvironmentsVariables = &TapaError{Code: 4100, Message: "Failed fetching environment's variables -> "}
)
