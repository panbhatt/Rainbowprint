package cmd

import (
	"flag"
	"fmt"

	internal "github.com/panbhatt/Rainbowprint/internal"
)

func parametersAndFlags() (bool, bool, []string) {

	help :=flag.Bool("help", false, "Display Help")
	helpShort :=flag.Bool("h", false, "Display Help")

	flag.Usage = printHelp

	flag.Parse()

	return *help || *helpShort,  flag.NArg() > 0 , flag.Args()

}

const VERSION = "1.0.0"
const HELP_MESSAGE = `Rainbow Print %s
Concatenate FILE(S) or standard input to standard output. 
When no FILE is passed read Standard Input. 

Examples:
	rainbowprint fileOne fileTwo                       # Output fileOne and fileTwo Contents
	rainbowprint                                       # Copy Standard Input to standard output
	echo "My Message" | rainbowprint                   # Display "My Message"
	fortune | rainbowprint                             # Display a rainbow cookie

	if you need more help, found a bug or want to suggest some new features: https://github.com/panbhatt/Rainbowprint 
`

func printHelp() {

	internal.PrintWithScanner(fmt.Sprintf(HELP_MESSAGE, VERSION))
}