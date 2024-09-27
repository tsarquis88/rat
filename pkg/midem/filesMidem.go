package midem

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func CreateDirManagers(dirPath string, originDir string) []IDataBytesManager {
	filesInDir := GetFilesInDir(dirPath)

	var managers []IDataBytesManager
	for _, file := range filesInDir {
		pathWithDir := dirPath + "/" + file
		if IsDir(pathWithDir) {
			managers = append(managers, CreateDirManagers(pathWithDir, originDir)...)
		} else {
			managers = append(managers, NewDataBytesFileManager(pathWithDir, originDir))
		}
	}
	return managers
}

func MixFiles(inputFiles []string, outputFile string) {
	fmt.Printf("Mix: %s\n", inputFiles)

	var managers []IDataBytesManager
	for _, file := range inputFiles {
		if IsDir(file) {
			fmt.Printf("Folder will be mixed: %s\n", file)
			fileWithoutSlash := strings.TrimSuffix(file, "/")

			var originDir string
			if filepath.Dir(fileWithoutSlash) == file {
				originDir = ""
			} else {
				originDir = filepath.Dir(fileWithoutSlash) + "/"
			}
			managers = append(managers, CreateDirManagers(fileWithoutSlash, originDir)...)
		} else {
			fmt.Printf("File will be mixed: %s\n", file)
			managers = append(managers, NewDataBytesFileManager(file, filepath.Dir(file)+"/"))
		}
	}

	var filesOriginDir []MetadataInput
	for _, manager := range managers {
		filesOriginDir = append(filesOriginDir, MetadataInput{manager.Name(), manager.Origin()})
	}

	dumpData := Dump(Generate(filesOriginDir))
	dumpData = append(dumpData, NewMixer(managers).Mix()...)

	if FileExists(outputFile) {
		panic("Output file exists")
	}

	NewDataBytesDumper(outputFile).Dump(dumpData)

	fmt.Printf("Files mixed into: %s\n", outputFile)
}

func DemixFiles(filesList []string, outputFolder string) {
	for _, inputFile := range filesList {
		fmt.Printf("Demix: %s\n", inputFile)

		fileManager := NewDataBytesFileManager(inputFile, "")

		for _, demixData := range Demix(fileManager) {
			err := os.MkdirAll(outputFolder+filepath.Dir(demixData.Filename), 0755)
			if err != nil {
				panic(err)
			}

			fmt.Printf("Writing file %s\n", outputFolder+demixData.Filename)
			err = os.WriteFile(outputFolder+demixData.Filename, demixData.Data, os.FileMode(demixData.Mode))
			if err != nil {
				panic(err)
			}
		}
	}
}
