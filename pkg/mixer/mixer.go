package mixer

import (
	"github.com/tsarquis88/file_mixer/pkg/dataBytesManager"
)

type mixer struct {
	managers []dataBytesManager.IDataBytesManager
}

func NewMixer(managers []dataBytesManager.IDataBytesManager) mixer {
	return mixer{managers}
}

func remove(slice []dataBytesManager.IDataBytesManager, s int) []dataBytesManager.IDataBytesManager {
	if len(slice) == 1 {
		return []dataBytesManager.IDataBytesManager{}
	}
	return append(slice[:s], slice[s+1:]...)
}

func (mixer mixer) Mix() (mixedData []byte) {
	var data []byte

	for {
		for index, manager := range mixer.managers {
			newBytes, n := manager.Read(1)
			if n == 0 {
				mixer.managers = remove(mixer.managers, index)
				continue
			}
			data = append(data, newBytes...)
		}

		if len(mixer.managers) == 0 {
			break
		}
	}

	return data
}
