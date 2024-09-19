package helpers

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func FindHash(hash string, directory string) (bool, int, error) {
	var foundHash int
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			calculatedHash, err := Sum(path)
			if err != nil {
				return err
			}
			if strconv.Itoa(calculatedHash) == hash {
				foundHash = calculatedHash
				return fmt.Errorf("found")
			}
		}
		return nil
	})

	if err != nil && err.Error() == "found" {
		return true, foundHash, nil
	}
	return false, -1, nil
}