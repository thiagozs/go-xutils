package files

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type Delimiter int

// byte SLASH_N = 0x0A; //new line
// byte EBCDIC_CR = 0x0D; //carriage return
// byte EBCDIC_NL = 0x15; //next line
// byte EBCDIC_LF = 0x25; //line feed

const (
	SLASH_N Delimiter = iota
	EBCDIC_CP500_CR
	EBCDIC_CP500_NL
	EBCDIC_CP500_LF
)

func (d Delimiter) Value() byte {
	return [...]byte{0x0A, 0x0D, 0x15, 0x25}[d]
}

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
		// if there's any error (including not exists), return false
		return false
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
	var filesInfo []os.FileInfo

	// use os.ReadDir for slightly better performance and clarity
	var walk func(string) error
	walk = func(p string) error {
		entries, err := os.ReadDir(p)
		if err != nil {
			return err
		}
		for _, e := range entries {
			fullPath := filepath.Join(p, e.Name())
			if e.IsDir() {
				if err := walk(fullPath); err != nil {
					return err
				}
				continue
			}
			info, err := e.Info()
			if err != nil {
				return err
			}
			filesInfo = append(filesInfo, info)
		}
		return nil
	}

	if err := walk(dirPath); err != nil {
		return nil, err
	}
	return filesInfo, nil
}

// FileExists checks if a file exists.
func (f *Files) FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}

// RemoveAllDir removes all files and directories in a directory specified by dirPath.
func (f *Files) RemoveAllDir(dirPath string) error {
	return os.RemoveAll(dirPath)
}

// ReadFileLines reads and returns the contents of a file specified by filePath line by line as a slice of strings.
func (f *Files) ReadFileLines(filePath string, opts ...Delimiter) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	if len(opts) == 0 {
		opts = append(opts, SLASH_N)
	}

	var lines []string
	reader := bufio.NewReader(file)
	delimiter := opts[0].Value()

	for {
		line, err := reader.ReadString(delimiter)
		if err != nil {
			if err == io.EOF {
				if len(line) > 0 {
					lines = append(lines, strings.TrimSuffix(line, string(delimiter)))
				}
				break
			}
			return nil, err
		}
		lines = append(lines, strings.TrimSuffix(line, string(delimiter)))
	}

	return lines, nil
}

// ReadFileByLineBytes reads and returns the contents of a file specified by filePath line by line as a slice of byte slices.
func (f *Files) ReadFileLinesBytes(filePath string, opts ...Delimiter) ([][]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	if len(opts) == 0 {
		opts = append(opts, SLASH_N)
	}

	var lines [][]byte
	reader := bufio.NewReader(file)
	delimiter := opts[0].Value()

	for {
		line, err := reader.ReadBytes(delimiter)
		if err != nil {
			if err == io.EOF {
				if len(line) > 0 {
					lines = append(lines, bytes.TrimSuffix(line, []byte{delimiter}))
				}
				break
			}
			return nil, err
		}
		lines = append(lines, bytes.TrimSuffix(line, []byte{delimiter}))
	}

	return lines, nil
}

// CompareSize compares the sizes of two files and returns true if they are equal.
func (f *Files) CompareSize(file1, file2 string) (bool, error) {
	info1, err := os.Stat(file1)
	if err != nil {
		return false, err
	}

	info2, err := os.Stat(file2)
	if err != nil {
		return false, err
	}

	return info1.Size() == info2.Size(), nil
}

// FileSizeBytes returns the size of a file in bytes.
func (f *Files) FileSizeBytes(filePath string) (int64, error) {
	info, err := os.Stat(filePath)
	if err != nil {
		return 0, err
	}
	return info.Size(), nil
}

// FileSizeKB returns the size of a file in kilobytes.
func (f *Files) FileSizeKB(filePath string) (int64, error) {
	info, err := os.Stat(filePath)
	if err != nil {
		return 0, err
	}
	return info.Size() / 1024, nil
}

// FileSizeMB returns the size of a file in megabytes.
func (f *Files) FileSizeMB(filePath string) (int64, error) {
	info, err := os.Stat(filePath)
	if err != nil {
		return 0, err
	}
	return info.Size() / 1024 / 1024, nil
}

// FileSizeGB returns the size of a file in gigabytes.
func (f *Files) FileSizeGB(filePath string) (int64, error) {
	info, err := os.Stat(filePath)
	if err != nil {
		return 0, err
	}
	return info.Size() / 1024 / 1024 / 1024, nil
}

// FileSizeTB returns the size of a file in terabytes.
func (f *Files) FileSizeTB(filePath string) (int64, error) {
	info, err := os.Stat(filePath)
	if err != nil {
		return 0, err
	}
	return info.Size() / 1024 / 1024 / 1024 / 1024, nil
}
