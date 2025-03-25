package errors

// ------------- Collection Repository
var (
	ErrCollectionsRetrieval      = &TapaError{Code: 3000, Message: "Failed fetching all collections \n"}
	ErrFoldersRetrieval          = &TapaError{Code: 3001, Message: "Failed fetching all folders \n"}
	ErrRequestSummariesRetrieval = &TapaError{Code: 3002, Message: "Failed fetching all request summaries \n"}
)
