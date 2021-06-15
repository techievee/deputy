package utilities

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

// GetCurrentPath Get the current path where the file is running
func GetCurrentPath() (string, error) {
	ex, err := os.Executable()
	if err != nil {
		return "", err
	}
	exPath := filepath.Dir(ex)
	return exPath, nil
}

// GetRootProjectPath Get the root path of the project where the tests are running
// Used for referencing the test files from the project_folder/test_samples
func GetRootProjectPath() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// Check if already in root directory (e.g. main.go)
	lastSlashIndex := strings.LastIndex(dir, "/")
	cName := dir[lastSlashIndex+1:]
	if cName == "deputy" {
		return dir
	}

	// Get parent directory
	parent := filepath.Dir(dir)
	lastSlashIndex = strings.LastIndex(parent, "/")
	pName := parent[lastSlashIndex+1:]

	// If not at root, continue getting parent
	for pName != "deputy" {
		parent = filepath.Dir(parent)
		lastSlashIndex = strings.LastIndex(parent, "/")
		pName = parent[lastSlashIndex+1:]
	}
	return parent
}
