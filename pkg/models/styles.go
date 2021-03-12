package models

// Styles struct is the main object for Styling the SimPDF.PDF document.
// This includes the TextSize, Color (of the text), TextVariant, LineSize,
// Border, BackgroundColor, and Font.
// Name will be used to select the style if created an stored in a list.
type Styles struct {
	Name            string
	TextSize        float64
	Color           RGBColor
	TextVariant     Variants
	LineSize        float64
	Border          Borders
	BackgroundColor RGBColor
	Font            Fonts
}
