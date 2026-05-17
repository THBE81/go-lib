package helpers

import (
	"os"
	"path/filepath"
	"regexp"
)

func ListFiles(dir string, pattern string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	regex, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}

	var files []string
	for _, entry := range entries {
		if !entry.IsDir() && regex.MatchString(entry.Name()) {
			files = append(files, entry.Name())
		}
	}
	return files, nil
}

func ForEachFiles(dir string, pattern string, callback func(fileName string)) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	regex, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if !entry.IsDir() && regex.MatchString(entry.Name()) {
			callback(entry.Name())
		}
	}
	return nil
}

func ChangeCurrentDir() {
	// Set the working directory to the project root
	dir, err := filepath.Abs("../") // Adjust this as needed
	if err != nil {
		panic(err)
	}

	err = os.Chdir(dir)
	if err != nil {
		panic(err)
	}
}
