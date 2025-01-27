package utils

import (
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

// FindFolderPath searches for a folder by name within a specific root directory
func FindFolderPath(root, folderName string, log *logrus.Logger) string {
	var foundPath string

	// Walk through the file system starting from the given root
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.WithError(err).Errorf("Error accessing path: %s", path)
			return nil
		}

		// Check if the folder name matches
		if info.IsDir() && info.Name() == folderName {
			foundPath = path
			log.Infof("Found folder: %s", path)
			return filepath.SkipDir // Stop searching once the folder is found
		}
		return nil
	})

	if err != nil {
		log.WithError(err).Error("Error walking through the filesystem")
	}

	return foundPath
}
