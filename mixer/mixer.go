package mixer

import (
	"fmt"
	"example.com/fileManager"
)

type mixer struct {
	managers []fileManager.FileManager
}

func New(files []string) (mixer) {
	var mixer = mixer {}
	for _, file := range files {
		fmt.Printf("Reading file: %s\n", file)
		mixer.managers = append(mixer.managers, fileManager.New(file, false))
	}
	fmt.Printf("Files read: %d\n", len(mixer.managers))
	return mixer
}

func remove(slice []fileManager.FileManager, s int) []fileManager.FileManager {
	if len(slice) == 1 {
		return []fileManager.FileManager{}
	}
    return append(slice[:s], slice[s+1:]...)
}

func (mixer mixer) Mix() {
	var data []byte

	for {
		for index, manager := range mixer.managers {
			newByte, n := manager.Read()
			if n == 0 {
				
				mixer.managers = remove(mixer.managers, index)
				continue
			}
			data = append(data, newByte)
		}
	
		if len(mixer.managers) == 0 {
			break
		}
	}

	fmt.Printf("Data: \n\n%s", data)
}

