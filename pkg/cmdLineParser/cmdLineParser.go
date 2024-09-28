package cmdLineParser

func remove(slice []string, s int) []string {
	if len(slice) == 1 {
		return []string{}
	}
	return append(slice[:s], slice[s+1:]...)
}

func Parse(args []string) (bool, []string) {
	argsQty := len(args)
	if argsQty < 2 {
		panic("Missing arguments")
	}
	args = remove(args, 0)

	rat := true
	for i, arg := range args {
		if arg == "-x" {
			rat = false
			args = remove(args, i)
			break
		}
	}

	if rat && len(args) < 2 {
		panic("Missing arguments")
	}

	return rat, args
}
