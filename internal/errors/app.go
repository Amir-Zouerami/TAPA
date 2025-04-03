package errors

// ------------- GENERAL APP ERRORS (1000)
var (
	ErrAppIconRead            = &TapaError{Code: 1000, Message: "App icon read error -> "}
	ErrAppConfigGeneration    = &TapaError{Code: 1001, Message: "App config generation error -> "}
	ErrEmbeddedFileRead       = &TapaError{Code: 1002, Message: "Embedded file read error -> "}
	ErrGetUserConfigDirectory = &TapaError{Code: 1003, Message: "Getting user config directory failed -> "}
	ErrCreateAppDirectory     = &TapaError{Code: 1004, Message: "Failed creating app directory -> "}
)

// ------------- DATABASE INITIALIZATION ERRORS (2000)
var (
	ErrSchemaRead          = &TapaError{Code: 2001, Message: "Database schema read error -> "}
	ErrSchemaCreation      = &TapaError{Code: 2002, Message: "Database schema creation error -> "}
	ErrOpeningDatabaseFile = &TapaError{Code: 2003, Message: "Opening database file failed -> "}
	ErrConnectingDatabase  = &TapaError{Code: 2004, Message: "Connecting to database failed -> "}
)
