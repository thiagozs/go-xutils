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

func TestRemoveFile(t *testing.T) {
	fm := New()
	testDir, err := os.MkdirTemp("", "testRemoveFile")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(testDir)

	testFile := filepath.Join(testDir, "testFile.txt")
	if err := os.WriteFile(testFile, []byte("test"), 0644); err != nil {
		t.Fatalf("Failed to write test file: %v", err)
	}

	if err := fm.RemoveFile(testFile); err != nil {
		t.Fatalf("Failed to remove file: %v", err)
	}

	if fm.IsFile(testFile) {
		t.Errorf("File should have been removed: %s", testFile)
	}
}

func TestRemoveDir(t *testing.T) {
	fm := New()
	testDir, err := os.MkdirTemp("", "testRemoveDir")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(testDir)

	testSubDir := filepath.Join(testDir, "subdir")
	if err := os.Mkdir(testSubDir, 0755); err != nil {
		t.Fatalf("Failed to create test subdir: %v", err)
	}

	if err := fm.RemoveDir(testSubDir); err != nil {
		t.Fatalf("Failed to remove directory: %v", err)
	}

	if fm.IsDirectory(testSubDir) {
		t.Errorf("Directory should have been removed: %s", testSubDir)
	}
}

func TestCreateDir(t *testing.T) {
	fm := New()
	testDir, err := os.MkdirTemp("", "testCreateDir")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(testDir)

	newDir := filepath.Join(testDir, "newdir")
	if err := fm.CreateDir(newDir); err != nil {
		t.Fatalf("Failed to create directory: %v", err)
	}

	if !fm.IsDirectory(newDir) {
		t.Errorf("Directory should have been created: %s", newDir)
	}
}

func TestCopyFile(t *testing.T) {
	fm := New()
	testDir, err := os.MkdirTemp("", "testCopyFile")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(testDir)

	srcFile := filepath.Join(testDir, "src.txt")
	dstFile := filepath.Join(testDir, "dst.txt")
	data := []byte("test data")
	if err := os.WriteFile(srcFile, data, 0644); err != nil {
		t.Fatalf("Failed to write source file: %v", err)
	}

	if err := fm.CopyFile(srcFile, dstFile); err != nil {
		t.Fatalf("Failed to copy file: %v", err)
	}

	dstData, err := os.ReadFile(dstFile)
	if err != nil {
		t.Fatalf("Failed to read destination file: %v", err)
	}

	if string(dstData) != string(data) {
		t.Errorf("Copied data does not match: %s", dstData)
	}
}

func TestCopyDir(t *testing.T) {
	fm := New()
	testDir, err := os.MkdirTemp("", "testCopyDir")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(testDir)

	srcDir := filepath.Join(testDir, "src")
	dstDir := filepath.Join(testDir, "dst")
	if err := os.Mkdir(srcDir, 0755); err != nil {
		t.Fatalf("Failed to create source directory: %v", err)
	}

	// Create a nested directory structure
	nestedDir := filepath.Join(srcDir, "nested")
	if err := os.Mkdir(nestedDir, 0755); err != nil {
		t.Fatalf("Failed to create nested directory: %v", err)
	}
	nestedFile := filepath.Join(nestedDir, "nested.txt")
	if err := os.WriteFile(nestedFile, []byte("nested data"), 0644); err != nil {
		t.Fatalf("Failed to write nested file: %v", err)
	}

	if err := fm.CopyDir(srcDir, dstDir); err != nil {
		t.Fatalf("Failed to copy directory: %v", err)
	}

	// Verify that the nested file was copied
	dstNestedFile := filepath.Join(dstDir, "nested", "nested.txt")
	dstData, err := os.ReadFile(dstNestedFile)
	if err != nil {
		t.Fatalf("Failed to read destination file: %v", err)
	}

	if string(dstData) != "nested data" {
		t.Errorf("Copied data does not match: %s", dstData)
	}
}

func TestMoveFile(t *testing.T) {
	fm := New()
	testDir, err := os.MkdirTemp("", "testMoveFile")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(testDir)

	srcFile := filepath.Join(testDir, "src.txt")
	dstFile := filepath.Join(testDir, "dst.txt")
	data := []byte("test data")
	if err := os.WriteFile(srcFile, data, 0644); err != nil {
		t.Fatalf("Failed to write source file: %v", err)
	}

	if err := fm.MoveFile(srcFile, dstFile); err != nil {
		t.Fatalf("Failed to move file: %v", err)
	}

	if fm.IsFile(srcFile) {
		t.Errorf("Source file should have been moved: %s", srcFile)
	}

	dstData, err := os.ReadFile(dstFile)
	if err != nil {
		t.Fatalf("Failed to read destination file: %v", err)
	}

	if string(dstData) != string(data) {
		t.Errorf("Moved data does not match: %s", dstData)
	}
}

func TestMoveDir(t *testing.T) {
	fm := New()
	testDir, err := os.MkdirTemp("", "testMoveDir")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(testDir)

	srcDir := filepath.Join(testDir, "src")
	dstDir := filepath.Join(testDir, "dst")
	if err := os.Mkdir(srcDir, 0755); err != nil {
		t.Fatalf("Failed to create source directory: %v", err)
	}

	// Create a nested directory structure
	nestedDir := filepath.Join(srcDir, "nested")
	if err := os.Mkdir(nestedDir, 0755); err != nil {
		t.Fatalf("Failed to create nested directory: %v", err)
	}
	nestedFile := filepath.Join(nestedDir, "nested.txt")
	if err := os.WriteFile(nestedFile, []byte("nested data"), 0644); err != nil {
		t.Fatalf("Failed to write nested file: %v", err)
	}

	if err := fm.MoveDir(srcDir, dstDir); err != nil {
		t.Fatalf("Failed to move directory: %v", err)
	}

	if fm.IsDirectory(srcDir) {
		t.Errorf("Source directory should have been moved: %s", srcDir)
	}

	// Verify that the nested file was moved
	dstNestedFile := filepath.Join(dstDir, "nested", "nested.txt")
	dstData, err := os.ReadFile(dstNestedFile)
	if err != nil {
		t.Fatalf("Failed to read destination file: %v", err)
	}

	if string(dstData) != "nested data" {
		t.Errorf("Moved data does not match: %s", dstData)
	}
}

func TestReadDir(t *testing.T) {
	fm := New()
	testDir, err := os.MkdirTemp("", "testReadDir")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(testDir)

	// Create a nested directory structure with files
	nestedDir := filepath.Join(testDir, "nested")
	if err := os.Mkdir(nestedDir, 0755); err != nil {
		t.Fatalf("Failed to create nested dir: %v", err)
	}
	file1 := filepath.Join(testDir, "file1.txt")
	if err := os.WriteFile(file1, []byte("content"), 0644); err != nil {
		t.Fatalf("Failed to write file1: %v", err)
	}

	// Read the directory
	files, err := fm.ReadDir(testDir)
	if err != nil {
		t.Fatalf("Failed to read directory: %v", err)
	}

	// Verify that only file1 is found and not files from nested directories
	if len(files) != 1 {
		t.Fatalf("Expected 1 file, found %d", len(files))
	}

	if files[0].Name() != "file1.txt" {
		t.Errorf("Expected file1.txt to be found, but found %s", files[0].Name())
	}
}

func TestCreateDirAll(t *testing.T) {
	fm := New()
	testDir, err := os.MkdirTemp("", "testCreateDirAll")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(testDir)

	newDir := filepath.Join(testDir, "newdir", "subdir")
	if err := fm.CreateDirAll(newDir); err != nil {
		t.Fatalf("Failed to create directory: %v", err)
	}

	if !fm.IsDirectory(newDir) {
		t.Errorf("Directory should have been created: %s", newDir)
	}
}

func TestRenameFile(t *testing.T) {
	fm := New()
	testDir, err := os.MkdirTemp("", "testRenameFile")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(testDir)

	srcFile := filepath.Join(testDir, "src.txt")
	dstFile := filepath.Join(testDir, "dst.txt")
	data := []byte("test data")
	if err := os.WriteFile(srcFile, data, 0644); err != nil {
		t.Fatalf("Failed to write source file: %v", err)
	}

	if err := fm.RenameFile(srcFile, dstFile); err != nil {
		t.Fatalf("Failed to rename file: %v", err)
	}

	if fm.IsFile(srcFile) {
		t.Errorf("Source file should have been renamed: %s", srcFile)
	}

	dstData, err := os.ReadFile(dstFile)
	if err != nil {
		t.Fatalf("Failed to read destination file: %v", err)
	}

	if string(dstData) != string(data) {
		t.Errorf("Renamed data does not match: %s", dstData)
	}
}

func TestFileExists(t *testing.T) {
	fm := New()
	testDir, err := os.MkdirTemp("", "testFileExists")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(testDir)

	testFile := filepath.Join(testDir, "testFile.txt")
	if err := os.WriteFile(testFile, []byte("test"), 0644); err != nil {
		t.Fatalf("Failed to write test file: %v", err)
	}

	if !fm.FileExists(testFile) {
		t.Errorf("File should exist: %s", testFile)
	}

	nonExistentFile := filepath.Join(testDir, "nonexistent.txt")
	if fm.FileExists(nonExistentFile) {
		t.Errorf("File should not exist: %s", nonExistentFile)
	}
}

func TestRemoveAllDir(t *testing.T) {
	fm := New()
	testDir, err := os.MkdirTemp("", "testRemoveAllDir")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(testDir)

	// Create a nested directory structure
	nestedDir := filepath.Join(testDir, "nested")
	if err := os.Mkdir(nestedDir, 0755); err != nil {
		t.Fatalf("Failed to create nested dir: %v", err)
	}
	nestedFile := filepath.Join(nestedDir, "nested.txt")
	if err := os.WriteFile(nestedFile, []byte("nested data"), 0644); err != nil {
		t.Fatalf("Failed to write nested file: %v", err)
	}

	if err := fm.RemoveAllDir(testDir); err != nil {
		t.Fatalf("Failed to remove directory: %v", err)
	}

	if fm.IsDirectory(testDir) {
		t.Errorf("Directory should have been removed: %s", testDir)
	}
}
