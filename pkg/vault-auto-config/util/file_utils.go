package util

import (
	"io"
	"os"
)

// Returns true if directory does not exist or is empty
func IsEmptyDir(dir string) (empty bool, err error) {
	_, err = os.Stat(dir)

	if os.IsNotExist(err) {
		empty = true
		err = nil
		return
	}

	if err != nil {
		return
	}

	f, err := os.Open(dir)

	if err != nil {
		return
	}

	defer func() {
		cerr := f.Close()
		if err == nil {
			err = cerr
		}
	}()

	_, err = f.Readdirnames(1)
	if err == io.EOF {
		empty = true
		err = nil
		return
	}

	return
}
