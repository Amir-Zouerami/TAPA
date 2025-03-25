package errors

// ------------- GENERAL APP ERRORS (1000)
var (
	ErrAppIconRead            = &TapaError{Code: 1000, Message: "App icon read error \n"}
	ErrAppConfigGeneration    = &TapaError{Code: 1001, Message: "App config generation error \n"}
	ErrEmbeddedFileRead       = &TapaError{Code: 1002, Message: "Embedded file read error \n"}
	ErrGetUserConfigDirectory = &TapaError{Code: 1003, Message: "Getting user config directory failed \n"}
	ErrCreateAppDirectory     = &TapaError{Code: 1004, Message: "Failed creating app directory \n"}
)

// ------------- DATABASE INITIALIZATION ERRORS (2000)
var (
	ErrSchemaRead          = &TapaError{Code: 2001, Message: "Database schema read error \n"}
	ErrSchemaCreation      = &TapaError{Code: 2002, Message: "Database schema creation error \n"}
	ErrOpeningDatabaseFile = &TapaError{Code: 2003, Message: "Opening database file failed \n"}
	ErrConnectingDatabase  = &TapaError{Code: 2004, Message: "Connecting to database failed \n"}
)
