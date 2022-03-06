package cmd

import (
	

	 "github.com/panbhatt/Rainbowprint/internal"

	 colorable "github.com/mattn/go-colorable"
)

func Main() {

	// This is done for windows support. 
	defer colorable.EnableColorsStdout(nil)()

	help, fileArguments, files := parametersAndFlags()


	if help {
		printHelp()
		return
	}

	if fileArguments {
		text , err := internal.JoinFiles(files)

		if(err != nil) {
			internal.PrintWithScanner(err.Error())
			return
		}

		internal.PrintWithScanner(text)
		return
	}

	internal.StartProcessFromStdin()

}