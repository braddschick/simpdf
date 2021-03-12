package models

// BorderWidths struct allows for the notation of border widths around any object
// for inclusion in a models.Styles
type BorderWidths struct {
	Left   float64
	Top    float64
	Right  float64
	Bottom float64
	Unit   Units
}

// Borders struct is the object for use in models.Styles
type Borders struct {
	Color RGBColor
	Width BorderWidths
}
