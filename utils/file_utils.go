package utils

import (
	"io/fs"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

// IsValidFolder checks if the provided path is a valid and accessible folder
func IsValidFolder(folderPath string) bool {
	info, err := os.Stat(folderPath)
	if err != nil || !info.IsDir() {
		return false
	}
	return true
}

// ListFiles retrieves all files from a folder, including their sizes
func ListFiles(folder string, log *logrus.Logger) map[string]int64 {
	files := make(map[string]int64)

	err := filepath.Walk(folder, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			log.WithError(err).Errorf("Error reading file: %s", path)
			return nil
		}

		// Only process files, ignore directories
		if !info.IsDir() {
			files[info.Name()] = info.Size() // Store file name and size
		}
		return nil
	})

	if err != nil {
		log.WithError(err).Error("Error walking the folder")
	}

	return files
}

// CompareFileLists compares two file maps and returns missing or different files
func CompareFileLists(oldFiles, newFiles map[string]int64) (missingInNew, missingInOld []string) {
	// Files present in the old folder but missing or different in the new folder
	for fileName, size := range oldFiles {
		if newSize, exists := newFiles[fileName]; !exists || newSize != size {
			missingInNew = append(missingInNew, fileName)
		}
	}

	// Files present in the new folder but missing or different in the old folder
	for fileName, size := range newFiles {
		if oldSize, exists := oldFiles[fileName]; !exists || oldSize != size {
			missingInOld = append(missingInOld, fileName)
		}
	}

	return missingInNew, missingInOld
}
