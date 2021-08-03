package main

import (
  "fmt"
  "os"
  "strings"
	"github.com/gookit/color"
)

func main() {
  if len(os.Args) < 2 {
    fmt.Printf("hc\nUsage: hc <hexcolor>\n")
    return
  }

  hex := os.Args[1]
  rgb := color.HexToRgb(hex)
  rgbString := strings.Trim(strings.Replace(fmt.Sprint(rgb), " ", ",", -1), "[]")
  color.HEXStyle("eee", hex).Printf("   rgb(%s)    \n", rgbString)
}
