package internal 


import (
	"os"
	"fmt"
	"strings"
	"bufio"

	rprint "github.com/panbhatt/Rainbowprint/pkg"
)

func JoinFiles(files []string) (string, error) {

	finalText := ""


	for i,fl := range files {
		fileContent, err := os.ReadFile(fl)

		if(err != nil) {
			return "",err
		}
		finalText += fl + "\n" + string(fileContent)  + "---------------------------------------------\n"
	}

	return finalText, nil

}

func PrintWithScanner(text string) {
		scanner := bufio.NewScanner(strings.NewReader(text))

		for i:=0; scanner.Scan();i++ {
			rgb := rprint.NewRGB(i)
			fmt.Printf("\033[38;2;%d;%d;%dm%s\033[0m\n", rgb.Red, rgb.Green, rgb.Blue, scanner.Text())
		}
}

func StartProcessFromStdin() {

	reader := bufio.NewReader(os.Stdio)

	for i:=0; true; i++ {
			input, err := reader.ReadRune()

			if(err != nil) {
				PrintWithScanner(err.Error())
				break
			}

			rgb := rprint.NewRGB(i)
			fmt.Printf("\033[38;2;%d;%d;%dm%s\033[0m\n", rgb.Red, rgb.Green, rgb.Blue, string(input))
	}


}