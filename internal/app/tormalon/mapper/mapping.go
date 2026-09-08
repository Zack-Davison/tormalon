package mapper

import "errors"

func MapArgumentsAndFlags(args []string) (map[string]string, error) {
	//create a map of each argument passed in
	//should be passed in the list of arguments without path

	if len(args) < 2 {
		return nil, errors.New("Not enough arguments provided")
	}

	argumentMap := map[string]string{
		"operation": args[0],
		"handler":   args[1],
	}

	//if args.

	//if lengthOfargument list is less than 2 throw an error

	return argumentMap, nil
}
