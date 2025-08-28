package main

import (
	"fmt"
	"net/url"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Units for each category (MUST match names used in convert())
var units = map[string][]string{
	"Weight":           {"Kilogram", "Gram", "Milligram"},
	"Length":           {"Meter", "Centimeter", "Millimeter"},
	"Liquid":           {"Liter", "Milliliter"},
	"Time":             {"Hour", "Minute", "Second"},
	"Temperature":      {"Celsius", "Fahrenheit", "Kelvin"},
	"Area":             {"Square Meter", "Square Kilometer", "Square Centimeter", "Square Millimeter", "Hectare", "Acre"},
	"Energy":           {"Joule", "Kilojoule", "Calorie", "Kilocalorie", "Watt-hour", "Kilowatt-hour"},
	"Power":            {"Watt", "Kilowatt", "Horsepower"},
	"Electric Current": {"Ampere", "Milliampere", "Kiloampere"},
	"Data Storage":     {"Byte", "Kilobyte", "Megabyte", "Gigabyte", "Terabyte"},
}

func convert(category, from, to string, value float64) float64 {
	// Short-circuit identical units
	if from == to {
		return value
	}

	switch category {
	// ----------------- WEIGHT -----------------
	case "Weight":
		var grams float64
		switch from {
		case "Kilogram":
			grams = value * 1000
		case "Gram":
			grams = value
		case "Milligram":
			grams = value / 1000
		}
		switch to {
		case "Kilogram":
			return grams / 1000
		case "Gram":
			return grams
		case "Milligram":
			return grams * 1000
		}

	// ----------------- LENGTH -----------------
	case "Length":
		var meters float64
		switch from {
		case "Meter":
			meters = value
		case "Centimeter":
			meters = value / 100
		case "Millimeter":
			meters = value / 1000
		}
		switch to {
		case "Meter":
			return meters
		case "Centimeter":
			return meters * 100
		case "Millimeter":
			return meters * 1000
		}

	// ----------------- LIQUID -----------------
	case "Liquid":
		var liters float64
		switch from {
		case "Liter":
			liters = value
		case "Milliliter":
			liters = value / 1000
		}
		switch to {
		case "Liter":
			return liters
		case "Milliliter":
			return liters * 1000
		}

	// ----------------- TIME -----------------
	case "Time":
		var seconds float64
		switch from {
		case "Hour":
			seconds = value * 3600
		case "Minute":
			seconds = value * 60
		case "Second":
			seconds = value
		}
		switch to {
		case "Hour":
			return seconds / 3600
		case "Minute":
			return seconds / 60
		case "Second":
			return seconds
		}

	// ----------------- TEMPERATURE -----------------
	case "Temperature":
		switch from {
		case "Celsius":
			if to == "Fahrenheit" {
				return value*9/5 + 32
			} else if to == "Kelvin" {
				return value + 273.15
			}
		case "Fahrenheit":
			if to == "Celsius" {
				return (value - 32) * 5 / 9
			} else if to == "Kelvin" {
				return (value-32)*5/9 + 273.15
			}
		case "Kelvin":
			if to == "Celsius" {
				return value - 273.15
			} else if to == "Fahrenheit" {
				return (value-273.15)*9/5 + 32
			}
		}

	// ----------------- AREA -----------------
	case "Area":
		var sqm float64
		switch from {
		case "Square Meter":
			sqm = value
		case "Square Kilometer":
			sqm = value * 1e6
		case "Square Centimeter":
			sqm = value / 10000
		case "Square Millimeter":
			sqm = value / 1e6
		case "Hectare":
			sqm = value * 10000
		case "Acre":
			sqm = value * 4046.8564224
		}
		switch to {
		case "Square Meter":
			return sqm
		case "Square Kilometer":
			return sqm / 1e6
		case "Square Centimeter":
			return sqm * 10000
		case "Square Millimeter":
			return sqm * 1e6
		case "Hectare":
			return sqm / 10000
		case "Acre":
			return sqm / 4046.8564224
		}

	// ----------------- ENERGY -----------------
	case "Energy":
		var joules float64
		switch from {
		case "Joule":
			joules = value
		case "Kilojoule":
			joules = value * 1000
		case "Calorie":
			joules = value * 4.184
		case "Kilocalorie":
			joules = value * 4184
		case "Watt-hour":
			joules = value * 3600
		case "Kilowatt-hour":
			joules = value * 3.6e6
		}
		switch to {
		case "Joule":
			return joules
		case "Kilojoule":
			return joules / 1000
		case "Calorie":
			return joules / 4.184
		case "Kilocalorie":
			return joules / 4184
		case "Watt-hour":
			return joules / 3600
		case "Kilowatt-hour":
			return joules / 3.6e6
		}

	// ----------------- POWER -----------------
	case "Power":
		var watts float64
		switch from {
		case "Watt":
			watts = value
		case "Kilowatt":
			watts = value * 1000
		case "Horsepower":
			watts = value * 745.7
		}
		switch to {
		case "Watt":
			return watts
		case "Kilowatt":
			return watts / 1000
		case "Horsepower":
			return watts / 745.7
		}

	// ----------------- ELECTRIC CURRENT -----------------
	case "Electric Current":
		var amps float64
		switch from {
		case "Ampere":
			amps = value
		case "Milliampere":
			amps = value / 1000
		case "Kiloampere":
			amps = value * 1000
		}
		switch to {
		case "Ampere":
			return amps
		case "Milliampere":
			return amps * 1000
		case "Kiloampere":
			return amps / 1000
		}

	// ----------------- DATA STORAGE -----------------
	case "Data Storage":
		var bytes float64
		switch from {
		case "Byte":
			bytes = value
		case "Kilobyte":
			bytes = value * 1024
		case "Megabyte":
			bytes = value * 1024 * 1024
		case "Gigabyte":
			bytes = value * 1024 * 1024 * 1024
		case "Terabyte":
			bytes = value * 1024 * 1024 * 1024 * 1024
		}
		switch to {
		case "Byte":
			return bytes
		case "Kilobyte":
			return bytes / 1024
		case "Megabyte":
			return bytes / (1024 * 1024)
		case "Gigabyte":
			return bytes / (1024 * 1024 * 1024)
		case "Terabyte":
			return bytes / (1024 * 1024 * 1024 * 1024)
		}
	}
	return 0
}

func main() {
	a := app.New()
	w := a.NewWindow("Unit Converter")

	categorySelect := widget.NewSelect(
		[]string{"Weight", "Length", "Liquid", "Time", "Temperature", "Area", "Energy", "Power", "Electric Current", "Data Storage"},
		nil,
	)

	fromSelect := widget.NewSelect([]string{}, nil)
	toSelect := widget.NewSelect([]string{}, nil)

	// Two entries (each combo has its own input)
	fromEntry := widget.NewEntry()
	fromEntry.SetPlaceHolder("Enter value (from)")

	toEntry := widget.NewEntry()
	toEntry.SetPlaceHolder("Enter value (to)")

	// Prevent feedback loops during programmatic updates
	updating := false

	updateTo := func() {
		if updating {
			return
		}
		if categorySelect.Selected == "" || fromSelect.Selected == "" || toSelect.Selected == "" {
			return
		}
		val, err := strconv.ParseFloat(fromEntry.Text, 64)
		if err != nil {
			// If invalid, just clear other side
			updating = true
			toEntry.SetText("")
			updating = false
			return
		}
		out := convert(categorySelect.Selected, fromSelect.Selected, toSelect.Selected, val)
		updating = true
		// use %.6g to avoid long floats like 1.9999999997
		toEntry.SetText(fmt.Sprintf("%.6g", out))
		updating = false
	}

	updateFrom := func() {
		if updating {
			return
		}
		if categorySelect.Selected == "" || fromSelect.Selected == "" || toSelect.Selected == "" {
			return
		}
		val, err := strconv.ParseFloat(toEntry.Text, 64)
		if err != nil {
			updating = true
			fromEntry.SetText("")
			updating = false
			return
		}
		out := convert(categorySelect.Selected, toSelect.Selected, fromSelect.Selected, val)
		updating = true
		fromEntry.SetText(fmt.Sprintf("%.6g", out))
		updating = false
	}

	// Wire up changes
	fromEntry.OnChanged = func(_ string) { updateTo() }
	toEntry.OnChanged = func(_ string) { updateFrom() }

	fromSelect.OnChanged = func(_ string) { updateTo() }
	toSelect.OnChanged = func(_ string) { updateTo() }

	categorySelect.OnChanged = func(cat string) {
		opts := units[cat]
		fromSelect.Options = opts
		toSelect.Options = opts
		fromSelect.Refresh()
		toSelect.Refresh()

		// Pick sensible defaults if available
		if len(opts) > 0 {
			fromSelect.SetSelected(opts[0])
		}
		if len(opts) > 1 {
			toSelect.SetSelected(opts[1])
		} else if len(opts) == 1 {
			toSelect.SetSelected(opts[0])
		}

		// Trigger conversion on category change
		updateTo()
	}

	link, _ := url.Parse("https://www.linkedin.com/in/kanishka-me/")
	linkedIn := widget.NewHyperlink("Kanishka Meddegoda", link)

	footer := container.NewHBox(
		widget.NewLabel("Created by:"),
		linkedIn,
	)

	footer = container.NewCenter(footer)

	w.SetContent(container.NewVBox(
		widget.NewLabel("Select Category"),
		categorySelect,
		widget.NewSeparator(),
		widget.NewLabel("From"),
		fromSelect,
		fromEntry,
		widget.NewLabel("To"),
		toSelect,
		toEntry,
		widget.NewSeparator(),
		footer,
	))

	w.Resize(fyne.NewSize(420, 500))
	w.ShowAndRun()
}

// This code implements a unit converter application using the Fyne GUI toolkit.
// It allows users to convert between various units of measurement across different categories.
// The conversion logic is encapsulated in the convert function, which handles the conversion based on the selected category and units.
// The GUI consists of dropdowns for selecting categories and units, and text entries for inputting values.
