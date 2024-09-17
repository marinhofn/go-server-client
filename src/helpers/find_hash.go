package helpers

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func FindHash(hash string, directory string) (bool, error) {
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			calculatedHash, err := Sum(path)
			fmt.Print("\ncalculatedhash: ", calculatedHash)
			if err != nil {
				return err
			}
			if strconv.Itoa(calculatedHash) == hash {
				return fmt.Errorf("found")
			}
		}
		return nil
	})

	if err != nil && err.Error() == "found" {
		return true, nil
	}
	return false, nil
}
