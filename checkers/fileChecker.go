package checkers

import (
	"errors"
	"os"
)

func FileExists(fileName string) bool {
	_, err := os.Stat(fileName)

	return !errors.Is(err, os.ErrNotExist)
}
