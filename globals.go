package simpdf

import (
	"github.com/braddschick/simpdf/pkg/models"
)

var (
	Header = HeaderFooters{
		LeftContent:   models.Contents{Align: models.Alignments{Left: true}},
		RightContent:  models.Contents{Align: models.Alignments{Right: true}},
		CenterContent: models.Contents{Align: models.Alignments{Center: true}},
	}
	Footer = HeaderFooters{
		LeftContent:   models.Contents{Align: models.Alignments{Left: true}},
		RightContent:  models.Contents{Align: models.Alignments{Right: true}},
		CenterContent: models.Contents{Align: models.Alignments{Center: true}},
	}
)
