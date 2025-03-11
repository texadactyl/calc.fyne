package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"go/token"
	"go/types"
)

var expression string

func main() {
	myApp := app.New()
	title := "Calculator"
	myWindow := myApp.NewWindow(title)
	label := widget.NewLabel(title)
	content := container.NewVBox(label)
	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(label.MinSize().Width+60, label.MinSize().Height+40)) // Add some padding

	output := widget.NewLabel(" ")

	buttons := []string{
		"7", "8", "9", "/",
		"4", "5", "6", "*",
		"1", "2", "3", "-",
		"0", ".", "=", "+",
		"C",
	}

	grid := container.NewGridWithColumns(4)

	for _, btnLabel := range buttons {
		button := widget.NewButton(btnLabel, func(txt string) func() {
			return func() {
				if txt == "=" {
					result, err := eval(expression)
					if err == nil {
						expression = result
						output.SetText(expression)
					} else {
						output.SetText("Error")
					}
				} else if txt == "C" {
					expression = ""
					output.SetText(" ")
				} else {
					expression += txt
					output.SetText(expression)
				}
			}
		}(btnLabel))
		grid.Add(button)
	}

	myWindow.SetContent(container.NewVBox(output, grid))
	myWindow.ShowAndRun()
}

func eval(expression string) (string, error) {
	fs := token.NewFileSet()
	typeAndValue, err := types.Eval(fs, nil, token.NoPos, "1.0 * "+expression)
	if err != nil {
		return "", fmt.Errorf("evaluation error")
	}

	if typeAndValue.Value == nil {
		return "", fmt.Errorf("invalid expression")
	}

	//return constant.StringVal(typeAndValue.Value), nil
	return typeAndValue.Value.String(), nil
}
