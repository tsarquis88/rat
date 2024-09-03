package cmdLineParser

func Parse(args []string) (string, []string) {
	argsQty := len(args)
	if argsQty < 2 {
		panic("Missing arguments")
	}

	if argsQty == 2 {
		// Demix
		return args[1], nil
	}

	// Mix
	return args[1], args[2:]
}
