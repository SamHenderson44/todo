package fileio

import (
	"os"
	"testing"
)

func TestCreateFile(t *testing.T) {
	t.Run("Successfully creates file", func(t *testing.T) {

		testFileName := "testFile.json"
		_, err := CreateFile(testFileName)

		defer func() {
			if err := os.Remove(testFileName); err != nil {
				t.Fatalf("Failed to delete test file: %v", err)
			}
		}()

		if err != nil {
			t.Errorf("Expected  no err, got %v", err)
		}

	})

	t.Run("Throws error with invalid file", func(t *testing.T) {

		testFileName := "/fakePath/testFile.json"
		_, err := CreateFile(testFileName)
		want := CreateFileError

		if err.Error() != want {
			t.Errorf("Expected err, got %v ", err)
		}

		defer func() {
			if _, err := os.Stat(testFileName); !os.IsNotExist(err) {
				os.Remove(testFileName)
			}
		}()

	})
}

func TestReadFile(t *testing.T) {
	t.Run("reads data correctly", func(t *testing.T) {
		tempFile, err := os.CreateTemp("", "testFile.txt")
		if err != nil {
			t.Fatalf("Failed to create file: %v", err)
		}

		defer os.Remove(tempFile.Name())

		want := "temp file"
		_, err = tempFile.WriteString(want)
		if err != nil {
			t.Fatalf("Failed to write to file: %v", err)
		}

		tempFile.Close()
		tempFile, err = os.Open(tempFile.Name())
		if err != nil {
			t.Fatalf("Failed to reopen file: %v", err)
		}
		defer tempFile.Close()

		got, err := ReadFile(tempFile)
		if err != nil {
			t.Errorf("ReadFile returned an error: %v", err)
		}

		if string(got) != want {
			t.Errorf("want %s, got %s", string(want), got)
		}

	})

	t.Run("errors if file doesn't exist", func(t *testing.T) {
		nonExistentFile, err := os.Open("non_existent_file.txt")
		if err == nil {
			t.Fatalf("Expected an error when opening a non-existent file, but got none")
		}

		_, err = ReadFile(nonExistentFile)
		if err == nil {
			t.Errorf("Expected an error when reading a non-existent file, but got none")
		}
	})
}
