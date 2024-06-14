package files

import (
	"bufio"
	"io"
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

// RemoveFile removes the file specified by filePath.
func (f *Files) RemoveFile(filePath string) error {
	return os.Remove(filePath)
}

// RemoveDir removes the directory specified by dirPath.
func (f *Files) RemoveDir(dirPath string) error {
	return os.RemoveAll(dirPath)
}

// CreateDir creates a new directory specified by dirPath.
func (f *Files) CreateDir(dirPath string) error {
	return os.Mkdir(dirPath, 0755)
}

// CreateDirAll creates a new directory specified by dirPath and all necessary parent directories.
func (f *Files) CreateDirAll(dirPath string) error {
	return os.MkdirAll(dirPath, 0755)
}

// RenameFile renames a file from oldPath to newPath.
func (f *Files) RenameFile(oldPath, newPath string) error {
	return os.Rename(oldPath, newPath)
}

// RenameDir renames a directory from oldPath to newPath.
func (f *Files) CopyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	return err
}

// CopyDir copies a directory from src to dst.
func (f *Files) CopyDir(src, dst string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}

		dstPath := filepath.Join(dst, relPath)
		if info.IsDir() {
			return os.Mkdir(dstPath, 0755)
		}

		return f.CopyFile(path, dstPath)
	})
}

// MoveFile moves a file from src to dst.
func (f *Files) MoveFile(src, dst string) error {
	return os.Rename(src, dst)
}

// MoveDir moves a directory from src to dst.
func (f *Files) MoveDir(src, dst string) error {
	return os.Rename(src, dst)
}

// ReadDir reads and returns all the entries in a directory specified by dirPath.
func (f *Files) ReadDir(dirPath string) ([]os.FileInfo, error) {
	var files []os.FileInfo
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() { // Optionally check if you want to add directories to the list
			files = append(files, info)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

// FileExists checks if a file exists.
func (f *Files) FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// RemoveAllDir removes all files and directories in a directory specified by dirPath.
func (f *Files) RemoveAllDir(dirPath string) error {
	return os.RemoveAll(dirPath)
}

// ReadFile reads and returns the entire content of a file specified by filePath as a string.
func (f *Files) ReadFile(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// ReadFileByLine reads and returns the contents of a file specified by filePath line by line as a slice of strings.
func (f *Files) ReadFileByLine(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string

	info, err := file.Stat()
	if err != nil {
		return nil, err
	}

	var maxSize int
	scanner := bufio.NewScanner(file)
	maxSize = int(info.Size())
	buffer := make([]byte, 0, maxSize)
	scanner.Buffer(buffer, maxSize)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
