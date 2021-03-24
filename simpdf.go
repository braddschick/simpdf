package simpdf

import (
	"time"

	"github.com/braddschick/simpdf/pkg/models"
	"github.com/jung-kurt/gofpdf"
)

// Doc is the interface used by the simpdf.SimPDF
type Doc interface {
	AddBottomLine(style models.Styles)
	AddImageCurrent(image Images)
	AddImageStandardPosition(image Images, stdPosition string)
	AddImageXY(image Images, x, y float64)
	AddMargins(margin models.Margins)
	AddNewLine()
	AddPageBreak()
	AddStyle(style []models.Styles)
	AddTable(tab Tables, altRowColor models.Styles, fixedWidth float64)
	AddTableHeader(table Tables, fixWidth float64)
	AddTableRows(table Tables, fixWidth float64)
	AppendStyle(style models.Styles)
	ChangeFont(style models.Styles)
	ChangePage(page models.Pages)
	CheckBottom() bool
	DistributeColumnsEvenly(numCols float64) float64
	DrawBottomLine(style models.Styles)
	Finish(fileOutput string)
	HeadingStart(styleType, text string)
	HeadingEnd(styleType string)
	PageHeight() float64
	PageWidth() float64
	Parser(style string, align models.Alignments, text string) string
	SetFont(fontFilePath string) error
	SetMargin(margin models.Margins)
	SetPage(pageType string, isLandscape bool)
	SetStyle(style models.Styles, fontOnly bool)
	StandardPosition(position string) (float64, float64)
	Start(pageType string, isLandscape bool, styles []models.Styles, margin models.Margins, customFontDirectory string)
	StringWidth(text string) float64
	StyleName(name string) (models.Styles, error)
	TableColumnWidth(table Tables) []float64
	WriteCenter(styleType, align models.Alignments, text string)
	Write(styleType, align models.Alignments, text string)
}

// SimPDF struct is the main object for the Simple PDF package.
type SimPDF struct {
	// Page denotes the size and orientation of the page(s) for the gofpdf.Pdf
	Page models.Pages
	// Style contains the various styles that are available for use in the SimPDF.PDF
	Style []models.Styles
	// Margin sets the margins for the models.Pages
	Margin models.Margins
	// PDF main PDF document from gofpdf.Pdf
	PDF gofpdf.Pdf
	// Title of the PDF document. Can be used as a variable for header/footer
	Title string
	// Author of the PDF document. Can be used as a variable for header/footer
	Author string
	// Keywords of the PDF document
	Keywords string
	// Subject for the PDF document
	Subject string
	// CreationDate gets authomatically set by running
	CreationDate time.Time
}
