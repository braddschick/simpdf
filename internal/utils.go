package internal

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// ValidateFilePath ensures the given filePath is accessible and is not a directory.
//
// Returns true is it is accessible, false if does not exist OR path is a directory.
func ValidateFilePath(filePath string) bool {
	fileInfo, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	}
	if fileInfo.IsDir() {
		return false
	}

	return true
}

// FileExtension returns the file extension of the given filePath in UPPERCASE for gofpdf.Pdf use.
func FileExtension(filePath string) string {
	return strings.ToUpper(filepath.Ext(filePath))
}

// MoveFilePath safely moves file path "filePath" to the newFilePath.
//
// Returns false if an error has occured, true if it was a successful operation
func MoveFilePath(filePath string, newFilePath string) bool {
	err := os.Rename(filePath, newFilePath)
	if err != nil {
		log.Fatalf("MoveFilePath incurred an Error moving from %s to %s.", filePath, newFilePath)
		return false
	}
	return true
}
// IsNumber safely checks the txt variable to see if it is a number
func IsNumber(txt string) bool {
	_, err := strconv.Atoi(txt)
	if err != nil {
		return false
	}
	return true
}

// IfError handles an error in a common format for ease of programming.
func IfError(msg string, err error, isPanic bool) {
	// if err != nil {
	// 	log.Fatalf("Error Message: %s || Actual Error: %s\n", msg, err.Error())
	// } else {
	// 	log.Fatal(msg)
	// }
	// if isPanic {
	// 	panic(msg)
	// }
}
