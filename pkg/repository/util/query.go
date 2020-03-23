package util

// GenerateArguments creates a list of interface from the Query class
func GenerateArguments(arguments []string) []interface{} {
	args := make([]interface{}, len(arguments))
	for i, argument := range arguments {
		args[i] = argument
	}

	return args
}
