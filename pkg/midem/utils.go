package midem

import (
	"crypto/sha256"
	"errors"
	"io"
	"os"
)

func FileExists(filepath string) bool {
	_, err := os.Stat(filepath)
	return !errors.Is(err, os.ErrNotExist)
}

func IsDir(filePath string) bool {
	fi, err := os.Stat(filePath)
	if err != nil {
		panic(err)
	}
	return fi.Mode().IsDir()
}

func GetFilesInDir(dirPath string) []string {
	dirHandle, err := os.ReadDir(dirPath)
	if err != nil {
		panic(err)
	}
	var filesList []string
	for _, file := range dirHandle {
		filesList = append(filesList, file.Name())
	}
	return filesList
}

func HashFile(filepath string) []byte {
	fileHandle, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer fileHandle.Close()
	h := sha256.New()
	if _, err := io.Copy(h, fileHandle); err != nil {
		panic(err)
	}
	return h.Sum(nil)
}
