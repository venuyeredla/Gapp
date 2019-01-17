package main

import (
	"fmt"

	"github.com/otiai10/gosseract"
)

func main() {
	client := gosseract.NewClient()
	defer client.Close()
	client.SetImage("C:\\Work\\WorkEng\\ImgToTe1.png")
	text, _ := client.Text()
	fmt.Println(text)
	// Hello, World!
}
