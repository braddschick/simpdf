package models

// Units struct allows the notation of what units are being used in the SimPDF.PDF document.
// Currently Design "pt" is only utilized.
type Units struct {
	Metric   bool
	Imperial bool
	Design   bool
}
