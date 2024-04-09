package files

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSaveFile(t *testing.T) {
	fm := New()
	testDir, err := os.MkdirTemp("", "test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(testDir)

	filePath := filepath.Join(testDir, "test.txt")
	data := []byte("Hello, world!")

	err = fm.SaveFile(filePath, data)
	if err != nil {
		t.Fatalf("Failed to save file: %v", err)
	}

	readData, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	if string(readData) != string(data) {
		t.Errorf("Expected file content: %s, got: %s", data, readData)
	}
}

func TestOpenFile(t *testing.T) {
	fm := New()
	testDir, err := os.MkdirTemp("", "test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(testDir)

	filePath := filepath.Join(testDir, "test.txt")
	data := []byte("Test content")
	if err := os.WriteFile(filePath, data, 0644); err != nil {
		t.Fatalf("Failed to write test file: %v", err)
	}

	readData, err := fm.OpenFile(filePath)
	if err != nil {
		t.Fatalf("Failed to open file: %v", err)
	}

	if string(readData) != string(data) {
		t.Errorf("Expected file content: %s, got: %s", data, readData)
	}
}

func TestAppendFile(t *testing.T) {
	fm := New()
	testDir, err := os.MkdirTemp("", "test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(testDir)

	filePath := filepath.Join(testDir, "appendTest.txt")
	initialData := []byte("Initial data. ")
	appendData := []byte("Appended data.")

	// Create the file with initial data
	if err := fm.SaveFile(filePath, initialData); err != nil {
		t.Fatalf("Failed to save file: %v", err)
	}

	// Append data to the file
	if err := fm.AppendFile(filePath, appendData); err != nil {
		t.Fatalf("Failed to append file: %v", err)
	}

	// Read and check the file content
	expectedData := append(initialData, appendData...)
	readData, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	if string(readData) != string(expectedData) {
		t.Errorf("Expected file content: %s, got: %s", expectedData, readData)
	}
}

func TestDirectoryExist(t *testing.T) {
	fm := New()
	testDir, err := os.MkdirTemp("", "test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(testDir)

	if !fm.DirectoryExist(testDir) {
		t.Errorf("Directory should exist: %s", testDir)
	}

	nonExistentDir := filepath.Join(testDir, "nonexistent")
	if fm.DirectoryExist(nonExistentDir) {
		t.Errorf("Directory should not exist: %s", nonExistentDir)
	}
}

func TestIsDirectory(t *testing.T) {
	fm := New()
	testDir, err := os.MkdirTemp("", "test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(testDir)

	if !fm.IsDirectory(testDir) {
		t.Errorf("Path should be a directory: %s", testDir)
	}

	testFile := filepath.Join(testDir, "testFile.txt")
	if err := os.WriteFile(testFile, []byte("test"), 0644); err != nil {
		t.Fatalf("Failed to write test file: %v", err)
	}

	if fm.IsDirectory(testFile) {
		t.Errorf("Path should not be a directory: %s", testFile)
	}
}

func TestIsFile(t *testing.T) {
	fm := New()
	testDir, err := os.MkdirTemp("", "test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(testDir)

	testFile := filepath.Join(testDir, "testFile.txt")
	if err := os.WriteFile(testFile, []byte("test"), 0644); err != nil {
		t.Fatalf("Failed to write test file: %v", err)
	}

	if !fm.IsFile(testFile) {
		t.Errorf("Path should be a file: %s", testFile)
	}

	if fm.IsFile(testDir) {
		t.Errorf("Path should not be a file: %s", testDir)
	}
}

func TestFileSearch(t *testing.T) {
	fm := New()
	testDir, err := os.MkdirTemp("", "testFileSearch")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(testDir)

	// Create a nested directory structure with files
	nestedDir := filepath.Join(testDir, "nested")
	if err := os.Mkdir(nestedDir, 0755); err != nil {
		t.Fatalf("Failed to create nested dir: %v", err)
	}
	targetFile := "target.txt"
	file1 := filepath.Join(testDir, targetFile)
	file2 := filepath.Join(nestedDir, targetFile)

	// Create target files
	if err := os.WriteFile(file1, []byte("content"), 0644); err != nil {
		t.Fatalf("Failed to write file1: %v", err)
	}
	if err := os.WriteFile(file2, []byte("content"), 0644); err != nil {
		t.Fatalf("Failed to write file2: %v", err)
	}

	// Search for the target file
	filesFound, err := fm.FileSearch(testDir, targetFile)
	if err != nil {
		t.Fatalf("FileSearch error: %v", err)
	}

	// Verify that both instances of the target file are found
	expectedFiles := []string{file1, file2}
	if len(filesFound) != len(expectedFiles) {
		t.Fatalf("Expected %d files, found %d", len(expectedFiles), len(filesFound))
	}

	for _, expectedFile := range expectedFiles {
		found := false
		for _, foundFile := range filesFound {
			if foundFile == expectedFile {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("File %s was not found by FileSearch", expectedFile)
		}
	}
}

func TestDirSearch(t *testing.T) {
	fm := New()
	testDir, err := os.MkdirTemp("", "testDirSearch")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(testDir)

	// Create a nested directory structure
	targetDirName := "targetDir"
	targetDir1 := filepath.Join(testDir, targetDirName)
	if err := os.Mkdir(targetDir1, 0755); err != nil {
		t.Fatalf("Failed to create targetDir1: %v", err)
	}
	nestedDir := filepath.Join(testDir, "nested")
	if err := os.Mkdir(nestedDir, 0755); err != nil {
		t.Fatalf("Failed to create nested dir: %v", err)
	}
	targetDir2 := filepath.Join(nestedDir, targetDirName)
	if err := os.Mkdir(targetDir2, 0755); err != nil {
		t.Fatalf("Failed to create targetDir2: %v", err)
	}

	// Search for the target directory
	dirsFound, err := fm.DirSearch(testDir, targetDirName)
	if err != nil {
		t.Fatalf("DirSearch error: %v", err)
	}

	// Verify that both instances of the target directory are found
	expectedDirs := []string{targetDir1, targetDir2}
	if len(dirsFound) != len(expectedDirs) {
		t.Fatalf("Expected %d directories, found %d", len(expectedDirs), len(dirsFound))
	}

	for _, expectedDir := range expectedDirs {
		found := false
		for _, foundDir := range dirsFound {
			if foundDir == expectedDir {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Directory %s was not found by DirSearch", expectedDir)
		}
	}
}
