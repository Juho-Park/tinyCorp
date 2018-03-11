package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"gopkg.in/headzoo/surf.v1"
)

func main() {
	args := os.Args
	for _, arg := range args {
		if strings.Contains(arg, "youtube") {
			ygom(arg)
		}
	}
	fmt.Println("ygom called")

	//ygom(args)
	//surfDemo()
}

func surfDemo() {
	bow := surf.NewBrowser()
	err := bow.Open("https://www.youtube.com")
	if err != nil {
		panic(err)
	}

	fmt.Println(bow.Title())

}

func ygom(url string) {
	//"youtube-dl -x --audio-format mp3 --audio-quality 0 "
	//"bluetooth-sendto --device=94:76:B7:00:71:B2 "

	yCmd := exec.Command("youtube-dl", url)

	opts := options{}

	yCmd.Args = append(yCmd.Args, "-x --audio-format mp3 --audio-quality 0")

	out := yCmd.Start()

	fmt.Println(out)
}
