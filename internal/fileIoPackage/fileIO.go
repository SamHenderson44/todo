package fileio

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
)

const CreateFileError = "error creating file"

func WriteToFile(file *os.File, toDos []byte) {
	permissions := 0644
	err := os.WriteFile(file.Name(), toDos, fs.FileMode(permissions))
	if err != nil {
		fmt.Println("Error encoding file: ", err)
		return
	}
}

func CreateFile(fileName string) (*os.File, error) {
	file, err := os.Create(fileName)
	if err != nil {
		return nil, errors.New(CreateFileError)
	}

	return file, nil
}

func ReadFile(file *os.File) ([]byte, error) {
	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}
	return bytes, nil
}
