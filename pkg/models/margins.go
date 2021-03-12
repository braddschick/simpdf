package models

// Margins struct allows the simple notation of margins for the SimPDF.PDF page
type Margins struct {
	Left   float64
	Top    float64
	Right  float64
	Bottom float64
	Unit   Units
}
