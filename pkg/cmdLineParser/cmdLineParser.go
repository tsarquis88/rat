package cmdLineParser

import "strconv"

type Parameters struct {
	Rat            bool
	List           bool
	OutputFolder   string
	OutputFile     string
	InputFiles     []string
	BlockingFactor uint
}

func remove(slice []string, s int) []string {
	if len(slice) == 1 {
		return []string{}
	}
	return append(slice[:s], slice[s+1:]...)
}

func Parse(args []string) Parameters {
	argsQty := len(args)
	if argsQty < 2 {
		panic("Missing arguments")
	}
	args = remove(args, 0)

	var params Parameters
	params.Rat = true
	params.List = false
	params.OutputFolder = ""
	params.BlockingFactor = 1
	for i, arg := range args {
		if arg == "-x" {
			params.Rat = false
			args = remove(args, i)
		} else if arg == "-C" {
			params.OutputFolder = args[i+1]
			args = remove(args, i)
			args = remove(args, i)
		} else if arg == "-t" {
			params.List = true
		} else if arg == "-b" {
			parsedSize, err := strconv.Atoi(args[i+1])
			if err != nil {
				panic(err)
			}
			params.BlockingFactor = uint(parsedSize)
			args = remove(args, i)
			args = remove(args, i)
		}
	}

	if !params.Rat && params.List {
		panic("Contradictory arguments: -x and -t")
	}

	if params.Rat && len(args) < 2 {
		panic("Missing arguments")
	}

	params.OutputFile = ""
	if params.Rat {
		params.OutputFile = args[0]
		args = remove(args, 0)
	}

	params.InputFiles = args
	return params
}
