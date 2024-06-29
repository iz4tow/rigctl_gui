package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"time"
	"os/exec"
	"strings"
	"strconv"
)

func getmode () (string) {
	getmodecmd := exec.Command("rigctl","-m2", "m")
	retmod, _ := getmodecmd.Output()
	mode := strings.SplitN(string(retmod), "\n", 2)
	return mode[0]
}

func updateButtons(AMButton,FMButton,USBButton,LSBButton *widget.Button) (){ //func to update buttons status
	//enable only active button
	AMButton.Disable()
	FMButton.Disable()
	USBButton.Disable()
	LSBButton.Disable()
	mode:=getmode()
	if mode == "AM" {
		FMButton.Enable()
		USBButton.Enable()
		LSBButton.Enable()
	}else if mode == "FM" {
		AMButton.Enable()
		USBButton.Enable()
		LSBButton.Enable()
	}else if mode == "USB" {
		FMButton.Enable()
		AMButton.Enable()
		LSBButton.Enable()
	}else if mode == "LSB" {
		FMButton.Enable()
		USBButton.Enable()
		AMButton.Enable()
	}
}

func main() {
	myApp := app.New()
	w := myApp.NewWindow("RadioManager")
	w.Resize(fyne.NewSize(300, 100))
	// Input field for frq
	input := widget.NewEntry()
	getfreqcmd := exec.Command("rigctl","-m2", "f")
	getpwrcmd := exec.Command("rigctl","-m2", "l", "RFPOWER")
	freq, err := getfreqcmd.Output()
	pwr,_ := getpwrcmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	input.SetPlaceHolder(string(freq))
	fmt.Println(pwr)

	//Buttons
	AMButton := widget.NewButton("AM", func() {
		fmt.Println("AM")
	})
	
	FMButton := widget.NewButton("FM", func() {
		fmt.Println("FM")
	})
	USBButton := widget.NewButton("USB", func() {
		fmt.Println("USB")
	})
	
	LSBButton := widget.NewButton("LSB", func() {
		fmt.Println("LSB")
	})
	
	
	go func() {
		for range time.Tick(time.Second) {
			updateButtons(AMButton,FMButton,USBButton,LSBButton)
		}
	}()


	frqButton := widget.NewButton("Set frq", func() { //Set frq
		_ = exec.Command("rigctl","-m2", "F",input.Text)
	})
	input.OnSubmitted = func(string) { //Set freq if ENTER is pressed inside input entry
		_ = exec.Command("rigctl","-m2", "F",input.Text)
	}

	valueLabel := widget.NewLabel("Current value: 0.00")
	mySlider := widget.NewSlider(0, 100)

	// Update the label text whenever the slider value changes
	mySlider.OnChanged = func(value float64) {
		valueLabel.SetText(fmt.Sprintf("RF Power: %.0f", value))
		power:=strconv.FormatFloat(value/100, 'f', -1, 64)
		_ = exec.Command("rigctl","-m2", "L", "RFPOWER", power)
	}
	// Arrange elements
	w.SetContent(
		container.NewBorder(
		container.NewBorder(
			nil,
			nil,
			nil,
			frqButton,
			input,
			),
		container.NewBorder(
			mySlider,
			valueLabel,
			nil,
			nil,
			),
			container.NewBorder(
				nil,
				nil,
				AMButton,
				FMButton,
			),
			container.NewBorder(
				nil,
				nil,
				USBButton,
				LSBButton,
			),
		),
	)
	






	w.ShowAndRun()
}

