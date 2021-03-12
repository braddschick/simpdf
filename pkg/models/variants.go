package models

// Variants struct allows the styling of text through Bold, Italic, Underline,
// or StrikeOut.
type Variants struct {
	Italic    bool
	Bold      bool
	Underline bool
	StrikeOut bool
}

// ToPDF function returns the proper gofpdf.PDF character for text variants.
func (v *Variants) ToPDF() string {
	if v.Italic {
		return "I"
	}
	if v.Bold {
		return "B"
	}
	if v.Underline {
		return "U"
	}
	if v.StrikeOut {
		return "S"
	}
	return ""
}
