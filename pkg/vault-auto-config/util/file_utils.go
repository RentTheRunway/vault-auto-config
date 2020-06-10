package util

import (
	"io"
	"os"
)

// Returns true if directory does not exist or is empty
func IsEmptyDir(dir string) (bool, error) {
	_, err := os.Stat(dir)

	if os.IsNotExist(err) {
		return true, nil
	}

	if err != nil {
		return false, err
	}

	f, err := os.Open(dir)

	if err != nil {
		return false, err
	}

	defer f.Close()

	_, err = f.Readdirnames(1)
	if err == io.EOF {
		return true, nil
	}

	return false, nil
}
