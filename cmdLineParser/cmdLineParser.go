package cmdLineParser

import "errors"

func Parse(args []string) ([]string, error) {
	if (len(args) <= 1) {
		return nil, errors.New("missing files")
	}

	listOfFiles := args[1:]
	return listOfFiles, nil
}
