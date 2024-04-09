package files

import (
	"os"
	"path/filepath"
)

type Files struct{}

func New() *Files {
	return &Files{}
}

// SaveFile writes data to a file specified by filePath. If the file doesn't exist, it is created. If it exists, it's overwritten.
func (f *Files) SaveFile(filePath string, data []byte) error {
	return os.WriteFile(filePath, data, 0644)
}

// OpenFile reads and returns the contents of the file specified by filePath.
func (f *Files) OpenFile(filePath string) ([]byte, error) {
	return os.ReadFile(filePath)
}

// AppendFile appends data to a file specified by filePath. If the file doesn't exist, it's created.
func (f *Files) AppendFile(filePath string, data []byte) error {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	return err
}

// DirectoryExist checks if a specified directory exists.
func (f *Files) DirectoryExist(dirPath string) bool {
	info, err := os.Stat(dirPath)
	if err != nil {
		return os.IsExist(err)
	}
	return info.IsDir()
}

// IsDirectory checks if the specified path is a directory.
func (f *Files) IsDirectory(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

// IsFile checks if the specified path is a file.
func (f *Files) IsFile(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

// FileSearch searches for files recursively starting from rootDir and returns a slice of file paths that match the provided fileName.
func (f *Files) FileSearch(rootDir, fileName string) ([]string, error) {
	var filesFound []string
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err // return error to stop the walk
		}
		if !info.IsDir() && info.Name() == fileName {
			filesFound = append(filesFound, path)
		}
		return nil
	})

	return filesFound, err
}

// DirSearch searches for directories recursively starting from rootDir and returns a slice of directory paths that match the provided dirName.
func (f *Files) DirSearch(rootDir, dirName string) ([]string, error) {
	var dirsFound []string
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err // return error to stop the walk
		}
		if info.IsDir() && info.Name() == dirName {
			dirsFound = append(dirsFound, path)
		}
		return nil
	})

	return dirsFound, err
}
