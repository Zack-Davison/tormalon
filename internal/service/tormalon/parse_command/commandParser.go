package parsecommand

import (
	"fmt"

	"github.com/zackdavison/tormalon/internal/app/tormalon/mapper"
)

// take in the command line arguments, mapping to, output them to a plan/configuration -> json plan -> then we can map that output to the input values for the template workflows

func ParseCommand(args []string) (string, error) {

	/*Function that will parse commands and route them to their designated handlers*/
	/*Handlers will be decided based on the second input past operation keyword e.g. tormalon create <handler> handler= java - dictates what*/

	argMap, err := mapper.MapArgumentsAndFlags(args)

	if err != nil {
		return "", fmt.Errorf("Failed to map arguments: %w", err)
	}

	fmt.Print(argMap)
	return "", nil
}
