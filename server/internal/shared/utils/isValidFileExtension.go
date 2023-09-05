package utils

import (
	"path/filepath"
	"reqwizard/internal/shared/constants"
)

func IsValidFileExtension(filename string) bool {
	ext := filepath.Ext(filename)

	return constants.FILE_EXTENSIONS[ext]
}

func IsValidAvatarImageExtension(filename string) bool {
	ext := filepath.Ext(filename)

	return constants.AVATAR_IMAGE_EXTENSIONS[ext]
}