package asciiART

import (
	"fmt"
	"os"
)

func PrintError(errorNum int) {
	switch errorNum {
	case 0:
		fmt.Println(`Usage: go run . [OPTION] [STRING] [BANNER]
EX: go run . --color=<color> --output=<outputfile.txt> <letters to be colored> "something" shadow`)
		os.Exit(1)
	case 1:
		fmt.Println(`Usage: go run . [OPTION] [STRING] [BANNER]
EX: go run . --align=<right> "something" shadow`)
		os.Exit(1)
	}
}
