package cmdLineParser

func remove(slice []string, s int) []string {
	if len(slice) == 1 {
		return []string{}
	}
	return append(slice[:s], slice[s+1:]...)
}

func Parse(args []string) (bool, string, []string) {
	argsQty := len(args)
	if argsQty < 2 {
		panic("Missing arguments")
	}
	args = remove(args, 0)

	rat := true
	outputFolder := ""
	for i, arg := range args {
		if arg == "-x" {
			rat = false
			args = remove(args, i)
		} else if arg == "-C" {
			outputFolder = args[i+1]
			args = remove(args, i)
			args = remove(args, i)
		}
	}

	if rat && len(args) < 2 {
		panic("Missing arguments")
	}

	return rat, outputFolder, args
}
