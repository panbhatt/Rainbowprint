package cmd

import (
	"flag"

	 "github.com/panbhatt/Rainbowprint/internal"
)

func Main() {

	help, fileArguments, files := parametersAndFlags()


	if help {
		printHelp()
		return
	}

	if fileArguments {
		text , err := internal.JoinFiles(fileArguments)

		if(err != nil) {
			internal.PrintWithScanner(error.Error())
			return
		}

		internal.PrintWithScanner(text)
		return
	}

	internal.StartProcessFromStdin()

}