package utils

import "os"

func DeleteFile(filePath string) error {
	err := os.Remove(filePath)
	if err != nil {
		return err
	}

	return nil
}