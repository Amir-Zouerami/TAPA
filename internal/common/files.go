package common

import (
	"io/fs"
	"os"

	"github.com/Amir-Zouerami/TAPA/internal/errors"
)

func ReadEmbeddedFile(embed fs.FS, name string) ([]byte, error) {
	data, err := fs.ReadFile(embed, name)
	if err != nil {
		return nil, errors.Wrap(errors.ErrEmbeddedFileRead, err)
	}

	return data, nil
}

func IsInDevelopmentMode() bool {
	return os.Getenv("TAPA_ENV") == "development"
}
