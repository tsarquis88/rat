package rat

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"errors"
	"io"
	"math"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

func FileRead(filepath string) []byte {
	fi, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	var fileData []byte
	for {
		buf := make([]byte, 1024)
		n, err := fi.Read(buf)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err)
		}
		fileData = append(fileData, buf[:n]...)
	}
	return fileData
}

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

func GetFilesInDir(dirPath string, recursive bool, includeDir bool) []string {
	dirHandle, err := os.ReadDir(dirPath)
	if err != nil {
		panic(err)
	}
	var filesList []string
	if includeDir {
		filesList = append(filesList, dirPath)
	}

	for _, file := range dirHandle {
		fileWithFolder := filepath.Join(dirPath, file.Name())
		if IsDir(fileWithFolder) {
			if recursive {
				filesList = append(filesList, GetFilesInDir(fileWithFolder, true, includeDir)...)
			}
		} else {
			filesList = append(filesList, fileWithFolder)
		}
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

func GzipCompress(inputData []byte) []byte {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)

	_, err := gz.Write(inputData)
	if err != nil {
		panic(err)
	}

	if err = gz.Flush(); err != nil {
		panic(err)
	}

	if err = gz.Close(); err != nil {
		panic(err)
	}

	return b.Bytes()
}

func GzipDecompress(inputData []byte) []byte {
	b := bytes.NewBuffer(inputData)

	var r io.Reader
	r, err := gzip.NewReader(b)
	if err != nil {
		panic(err)
	}

	var resB bytes.Buffer
	_, err = resB.ReadFrom(r)
	if err != nil {
		panic(err)
	}

	return resB.Bytes()
}

func OctalToDecimal(octal []byte, len int) uint {
	decimal := uint(0)
	for i := len - 1; i >= 0; i-- {
		if octal[i] != 48 {
			exp := len - 1 - i
			value := uint(octal[i]-48) * uint(math.Pow(float64(8), float64(exp)))
			decimal = decimal + value
		}
	}
	return decimal
}

func DecimalToOctal(decimal uint) []byte {
	var remainders []byte
	lastDecimal := decimal
	for {
		quotient := lastDecimal / 8
		remainders = append(remainders, byte(lastDecimal-quotient*8+48))
		if quotient == 0 {
			break
		}
		lastDecimal = quotient
	}
	slices.Reverse(remainders)
	return remainders
}

func FillWith(origin []byte, value byte, length uint) []byte {
	res := origin
	for i := uint(len(origin)); i < length; i++ {
		res = append(res, value)
	}
	return res
}

func TrimPrefixRecursive(s string, prefix string) string {
	if !strings.HasPrefix(s, prefix) {
		return s
	}
	return TrimPrefixRecursive(s[len(prefix):], prefix)
}
