package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gookit/color"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("hc\nUsage: hc <hexcolor>\n")
		return
	}

	input := os.Args[1]
	if strings.HasPrefix(input, "#") {
		input = input[1:]
	}
	hin, _ := strconv.ParseInt(input, 16, 64)
	white, _ := strconv.ParseInt("ffffff", 16, 64)
	fg := fmt.Sprintf("%x", (white - hin))
	rgb := color.HexToRgb(input)
	rgbString := strings.Trim(strings.Replace(fmt.Sprint(rgb), " ", ",", -1), "[]")
	color.HEXStyle(fg, input).Printf("   rgb(%s)    \n", rgbString)
}
