package trash

import (
  "fmt"

  "fyne.io/fyne/v2/app"
  "fyne.io/fyne/v2/container"
  "fyne.io/fyne/v2/widget"
)

func main() {
  myWindow := app.New().NewWindow("Slider Example")

  // Create a label to display the current slider value
  valueLabel := widget.NewLabel("Current value: 0.00")

  // Create a slider with min and max values
  mySlider := widget.NewSlider(0, 100)

  // Update the label text whenever the slider value changes
  mySlider.OnChanged = func(value float64) {
    valueLabel.SetText(fmt.Sprintf("Current value: %.2f", value))
  }

  // Arrange the slider and label in a vertical stack
  content := container.NewVBox(mySlider, valueLabel)

  // Set the window content and show it
  myWindow.SetContent(content)
  myWindow.ShowAndRun()
}

