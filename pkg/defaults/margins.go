package defaults

import "github.com/braddschick/simpdf/pkg/models"

var (
	// WordMargins returns the default margins from a MS Word Document
	// 1 inch / 2.54 centimeters
	Word_Margins = models.Margins{
		Left:   72.0,
		Top:    72.0,
		Right:  72.0,
		Bottom: 72.0,
		Unit:   models.Units{Design: true},
	}

	// ModerateMargins returns the default moderate margins from a MS Word Document
	// Top & Bottom 1 inch / 2.54 centimeters | Left & Right 3/4 inch / 1.905 centimeters
	Moderate_Margins = models.Margins{
		Left:   54.0,
		Top:    72.0,
		Right:  54.0,
		Bottom: 72.0,
		Unit:   models.Units{Design: true},
	}

	// NarrowMargins returns the narrow margins from a MS Word Document
	// 1/2 inch / 1.27 centimeters
	Narrow_Margins = models.Margins{
		Left:   36.0,
		Top:    36.0,
		Right:  36.0,
		Bottom: 36.0,
		Unit:   models.Units{Design: true},
	}

	// SuperNarrowMargins returns the narrowest of margins. Most likely
	// to cause printing issues depending on your printer.
	// 1/4 inch / 0.635 centimeters
	SuperNarrow_Margins = models.Margins{
		Left:   18.0,
		Top:    18.0,
		Right:  18.0,
		Bottom: 18.0,
		Unit:   models.Units{Design: true},
	}
)
