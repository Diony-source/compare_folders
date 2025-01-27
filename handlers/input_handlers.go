package handlers

import (
	"bufio"
	"compare_folders/services"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

// StartComparison prompts the user to enter folder paths and initiates the comparison process
func StartComparison(log *logrus.Logger) {
	reader := bufio.NewReader(os.Stdin)

	// Get the path to the old folder
	fmt.Print("Enter the path to the first folder (Old folder): ")
	oldFolderPath, _ := reader.ReadString('\n')
	oldFolderPath = sanitizeInput(oldFolderPath)

	// Get the path to the new folder
	fmt.Print("Enter the path to the second folder (New folder): ")
	newFolderPath, _ := reader.ReadString('\n')
	newFolderPath = sanitizeInput(newFolderPath)

	log.Infof("Received folder paths:\nOld Folder: %s\nNew Folder: %s", oldFolderPath, newFolderPath)

	// Pass the folder paths to the service layer for comparison
	services.CompareFolders(log, oldFolderPath, newFolderPath)
}

// sanitizeInput trims unnecessary spaces and newline characters from user input
func sanitizeInput(input string) string {
	return strings.TrimSpace(input)
}
