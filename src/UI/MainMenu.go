package UI

import (
	"fmt"
	"math"

	"github.com/fatih/color"
	"github.com/tecnologer/asciiart"
)

//MainMenu interface
type MainMenu struct {
	Title   string
	Options map[int]string
}

var options = map[int]string{
	1: "New",
	2: "View",
	3: "Modify",
	4: "Remove",
	5: "Exit",
}

const (
	headerChar   = "*"
	sidePadRight = 10
	sidePadLeft  = 10
	optionsLabel = "Main menu - Options"
)

//ShowMenu display the menu
func (mm MainMenu) ShowMenu() {
	color.Yellow(chars.ParsefBackground(mm.Title, '-', '|'))
	mm.PrintOptions()
}

//ShowMenuNHeader display the menu and default header with the title speicified
func (mm MainMenu) ShowMenuNHeader() {
	var header string
	var headerLen = len(optionsLabel)

	if len(mm.Title) > len(optionsLabel) {
		headerLen = len(mm.Title)
	}

	headerLen += sidePadLeft + sidePadRight
	for i := headerLen; i > 0; i-- {
		header += headerChar
	}
	padR, padL := getPads(mm.Title, header)
	color.Yellow(header)
	color.Yellow(fmt.Sprintf("*%s%s%s*\n", padR, mm.Title, padL))
	color.Yellow(header)

	padR, padL = getPads(optionsLabel, header)
	color.Yellow(fmt.Sprintf("*%s%s%s*\n", padR, optionsLabel, padL))
	color.Yellow(header)

	mm.PrintOptions()
}

func (mm MainMenu) PrintOptions() {
	if len(mm.Options) == 0 {
		mm.Options = options
	}

	for i := 1; i <= len(mm.Options); i++ {
		fmt.Printf("%d.- %s\n", i, mm.Options[i])
	}
}

//getPads return spaces in both sides
func getPads(word string, header string) (string, string) {
	var padR string
	var padL string
	var padRLen int
	var padLLen int
	var totalPads = int(math.Abs(float64(len(word) - len(header) + 2)))

	if (totalPads % 2) == 0 {
		padRLen = int(totalPads / 2)
		padLLen = padRLen
	} else {
		padRLen = int(totalPads/2) + 1
		padLLen = int(totalPads / 2)
	}

	for length := padRLen; length > 0; length-- {
		padR += " "
	}

	for length := padLLen; length > 0; length-- {
		padL += " "
	}

	return padR, padL
}
