package models

// Pages struct is the main page size for SimPDF.PDF
type Pages struct {
	Size        string
	Width       float64
	Height      float64
	IsLandscape bool
	Unit        Units
}

// ToPDFOrientation function returns the proper "L" for landscape orientation and
// "P" for portrait orientation. This is required for gofpdf.PDF
func (p *Pages) ToPDFOrientation() string {
	if p.IsLandscape {
		return "L"
	}
	return "P"
}
