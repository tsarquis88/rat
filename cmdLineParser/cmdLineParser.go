package cmdLineParser

import "errors"

func Parse(args []string) (string, []string, error) {
	if (len(args) <= 2) {
		return "", nil, errors.New("missing files")
	}

	listOfFiles := args[2:]
	return args[1], listOfFiles, nil
}
