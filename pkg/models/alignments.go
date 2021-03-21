package models

// Alignments struct simply Center, Left, and Right text alignment.
// All values are false by default. If no Alignment has been marked true
// then left "L" alignment is chosen.
type Alignments struct {
	Center  bool
	Left    bool
	Right   bool
	Justify bool
}

// ToPDF function returns the correct string character for the gofpdf.PDF
// document.
func (a *Alignments) ToPDF() string {
	if a.Center {
		return "C"
	}
	if a.Left {
		return "L"
	}
	if a.Right {
		return "R"
	}
	if a.Justify {
		return "J"
	}
	return "L"
}

// ToHTML function returns the correct string character for the gofpdf.PDF HTML
// style writing document.
func (a *Alignments) ToHTML() string {
	if a.Center {
		return "<center>$1</center>"
	}
	if a.Left {
		return "<left>$1</left>"
	}
	if a.Right {
		return "<right>$1</right>"
	}
	return "<left>$1</left>"
}
