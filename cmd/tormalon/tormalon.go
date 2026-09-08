package main

import (
	"fmt"
	"os"

	parsecommand "github.com/zackdavison/tormalon/internal/app/tormalon/parse_command"
)

func main() {
	argsWithoutPath := os.Args[1:]
	result, err := parsecommand.ParseCommand(argsWithoutPath)
	if err != nil {
		fmt.Print(err)
	}
	_ = result

}
