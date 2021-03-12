package defaults

import (
	"github.com/braddschick/simpdf/pkg/models"
)

var (
	// Letter size page 8.5 inches x 11 inches
	// All Sizes are in Design Units "points" pt
	Letter = models.Pages{
		Size:   "Letter",
		Width:  612,
		Height: 792,
		Unit:   models.Units{Design: true},
	}
	// Legal size page 8.5 inches x 14 inches
	// All Sizes are in Design Units "points" pt
	Legal = models.Pages{
		Size:   "Legal",
		Width:  612,
		Height: 1008,
		Unit:   models.Units{Design: true},
	}

	// Tabloid size page 11 inches x 14 inches
	// All Sizes are in Design Units "points" pt
	Tabloid = models.Pages{
		Size:   "Tabloid",
		Width:  792,
		Height: 1224,
		Unit:   models.Units{Design: true},
	}

	// Ledger size page which is very close to the Tabloid size
	// All Sizes are in Design Units "points" pt
	Ledger = models.Pages{
		Width:  791,
		Height: 1224.7,
		Unit:   models.Units{Design: true},
	}

	// AnsiA size page is the US Standard ANSI A
	// All Sizes are in Design Units "points" pt
	AnsiA = models.Pages{
		Width:  612.4,
		Height: 791,
		Unit:   models.Units{Design: true},
	}

	// A1 is the A1 page standard size
	// All Sizes are in Design Units "points" pt
	A1 = models.Pages{
		Width:  1684,
		Height: 2384.2,
		Unit:   models.Units{Design: true},
	}

	// A2 is the A2 page standard size
	// All Sizes are in Design Units "points" pt
	A2 = models.Pages{
		Width:  1190.7,
		Height: 1684,
		Unit:   models.Units{Design: true},
	}

	// A3 is the A3 standard page size
	// All Sizes are in Design Units "points" pt
	A3 = models.Pages{
		Size:   "A3",
		Width:  842,
		Height: 1190.7,
		Unit:   models.Units{Design: true},
	}

	// A4 is the A4 standard page size
	// All Sizes are in Design Units "points" pt
	A4 = models.Pages{
		Size:   "A4",
		Width:  595.4,
		Height: 842,
		Unit:   models.Units{Design: true},
	}

	// A5 is the A5 standard page size
	// All Sizes are in Design Units "points" pt
	A5 = models.Pages{
		Size:   "A5",
		Width:  419.6,
		Height: 595.4,
		Unit:   models.Units{Design: true},
	}
)
