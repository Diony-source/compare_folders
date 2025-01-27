package services

import (
	"compare_folders/utils"
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
)

// CompareFolders validates folder paths and compares their contents
func CompareFolders(log *logrus.Logger, oldFolder, newFolder string) {
	// Validate the old folder path
	if !utils.IsValidFolder(oldFolder) {
		log.Errorf("Invalid folder path: %s", oldFolder)
		fmt.Printf("Error: Invalid folder path '%s'\n", oldFolder)
		return
	}

	// Validate the new folder path
	if !utils.IsValidFolder(newFolder) {
		log.Errorf("Invalid folder path: %s", newFolder)
		fmt.Printf("Error: Invalid folder path '%s'\n", newFolder)
		return
	}

	// List files in both folders
	oldFiles := utils.ListFiles(oldFolder, log)
	newFiles := utils.ListFiles(newFolder, log)

	// Compare the file lists to find missing or different files
	missingInNew, missingInOld := utils.CompareFileLists(oldFiles, newFiles)

	// Print formatted results to console
	fmt.Println("=== Comparison Results ===")
	if len(missingInNew) > 0 {
		fmt.Printf("Files missing in New folder:\n%s\n", formatFileList(missingInNew))
		log.Infof("Files missing in New folder:\n%s", formatFileList(missingInNew))
	} else {
		fmt.Println("No files are missing in the New folder.")
	}

	if len(missingInOld) > 0 {
		fmt.Printf("Files missing in Old folder:\n%s\n", formatFileList(missingInOld))
		log.Infof("Files missing in Old folder:\n%s", formatFileList(missingInOld))
	} else {
		fmt.Println("No files are missing in the Old folder.")
	}
}

// formatFileList formats a list of file names into a readable string
func formatFileList(fileList []string) string {
	var formattedList strings.Builder
	for _, file := range fileList {
		formattedList.WriteString(fmt.Sprintf("- %s\n", file))
	}
	return formattedList.String()
}
