package simpdf

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/braddschick/simpdf/internal"
	"github.com/braddschick/simpdf/pkg/defaults"
	"github.com/braddschick/simpdf/pkg/models"
	"github.com/jung-kurt/gofpdf"
)

// AppendStyle allows the addition of more styles to the PDF document for usage by the developer.
func (s *SimPDF) AppendStyle(styles models.Styles) {
	s.Style = append(s.Style, styles)
}

// AddStyle replaces all styles with the supplied styles to the PDF document for usage by the developer.
func (s *SimPDF) AddStyle(styles []models.Styles) {
	s.Style = styles
}

// SetMargin sets the Margins for the PDF document.
// Use of a default is allowed here as defaults.WordMargins()
func (s *SimPDF) SetMargin(margin models.Margins) {
	s.Margin = margin
}

// SetFont sets the font to be used if it is not a standard font already provided.
// This will verify the file exists and if it is a directory the first TTF will be used.
func (s *SimPDF) SetFont(fontFilePath string) error {
	var font models.Fonts
	fileFont, err := os.Stat(fontFilePath)
	if os.IsNotExist(err) {
		log.Fatalf("Font file %s cannot be accessed or does not exist.", fontFilePath)
		panic(err)
	}
	if fileFont.IsDir() {
		log.Print("Font path given was a directory. The first ttf found will be used.")
		font.Directory = fontFilePath
		fonts, _ := filepath.Glob(fontFilePath + "*.ttf")
		if len(fonts) > 0 {
			// Need to remove path and only return the filename
			font.Name = fonts[0]
			font.IsValid = true
		} else {
			err := fmt.Sprintf("Font file(s) %s cannot be accessed or do not exist. No File with an extension of ttf exists.", fontFilePath)
			return errors.New(err)
		}
	} else {
		font.Directory = filepath.Dir(fontFilePath)
		font.Name = fileFont.Name()
		font.IsValid = true
	}
	return nil
}

// SetPage sets the page format for the PDF document. This should not be used directly but is here
// for context only. SimPDF.Start() should be used in lieu thereof.
func (s *SimPDF) SetPage(pageType string, isLandscape bool) {
	switch p := strings.ToLower(pageType); p {
	case "letter":
		s.Page = defaults.Letter
	case "legal":
		s.Page = defaults.Legal
	case "tabloid":
		s.Page = defaults.Tabloid
	case "ledger":
		s.Page = defaults.Ledger
	case "ansi":
		s.Page = defaults.AnsiA
	case "a1":
		s.Page = defaults.A1
	case "a2":
		s.Page = defaults.A2
	case "a3":
		s.Page = defaults.A3
	case "a4":
		s.Page = defaults.A4
	case "a5":
		s.Page = defaults.A5
	default:
		s.Page = defaults.Letter
	}
	s.Page.IsLandscape = isLandscape
}

// StyleName returns the style as directed by the name given. If the name given is not found
// then "Normal" will be looked for to use in lieu thereof.
func (s *SimPDF) StyleName(name string) (models.Styles, error) {
	var out models.Styles
	for _, t := range s.Style {
		if strings.ToLower(t.Name) == strings.ToLower(name) {
			out = t
			break
		}
	}
	if out.Name == "" {
		out, _ = s.StyleName("Normal")
	}
	if out.TextSize == 0 {
		err := fmt.Sprintf("Could not find the Style %s by name, and the Style 'Normal' does not exist.", name)
		return models.Styles{}, errors.New(err)
	}
	return out, nil
}

// AddNewLine Adds a new line to the PDF document the same line height as previously used.
func (s *SimPDF) AddNewLine(size float64) {
	if size == 0 {
		s.PDF.Ln(-1)
	} else {
		s.PDF.Ln(size)
	}
}

// DrawBottomLine Draws a simple bottom line under text as directed by the style given.
// This is useful for titles, headings, or other emphasis types of text. This is typically
// not directly called but rather used as part of SimPDF.AddBottomLine()
func (s *SimPDF) DrawBottomLine(style models.Styles) {
	_, y := s.PDF.GetXY()
	r, g, b := s.PDF.GetDrawColor()
	lw := s.PDF.GetLineWidth()
	y = y + style.LineSize + 2
	s.PDF.SetDrawColor(int(style.Border.Color.Red), int(style.Border.Color.Green), int(style.Border.Color.Blue))
	s.PDF.SetLineWidth(style.Border.Width.Bottom)
	xStart, _, xEnd, _ := s.PDF.GetMargins()
	s.PDF.Line(xStart+2, y, (s.Page.Width-xEnd)-2, y)
	s.PDF.SetXY(xStart, y+style.Border.Width.Bottom+5)
	s.PDF.SetLineWidth(lw)
	s.PDF.SetDrawColor(r, g, b)
}

// AddBottomLine is the manin funciton for adding a bottom line to the text as directed by the
// style given.
func (s *SimPDF) AddBottomLine(style models.Styles) {
	if style.Border.Width.Bottom > 0 {
		s.DrawBottomLine(style)
	}
}

// StandardPosition returns the X, Y coordinates of a standard position given in position.
// TL/C/R - Top Left/Center/Right
// CL/C/R - Center Left/Center/Right
// BL/C/R - Bottom Left/Center/Right
func (s *SimPDF) StandardPosition(position string) (float64, float64) {
	// position {string} can be anyone of these values
	// "tl", "tc", "tr"
	// "cl", "cc", "cr"
	// "bl", "bc", "br"
	var x, y float64
	idx := strings.Split(strings.ToLower(position), "")
	if len(idx) == 2 {
		switch p := idx[0]; p {
		case "t":
			y = s.Margin.Top
		case "c":
			y = s.PageHeight() / 2
		case "b":
			y = s.PageHeight() - s.Margin.Bottom
		}
		switch p := idx[1]; p {
		case "l":
			x = s.Margin.Left
		case "c":
			x = s.PageWidth() / 2
		case "r":
			x = s.PageWidth() - s.Margin.Right
		}
	} else {
		internal.IfError("Not enough was passed for Standard Position", nil, false)
	}
	return x, y
}

// AddPageBreak Simply adds a new page to the PDF document
func (s *SimPDF) AddPageBreak() {
	s.PDF.AddPage()
}

// PageHeight returns the current PDF document height depending on orientation Portrait or Landscape.
// Important note: this is the ENTIRE page and not with margins
func (s *SimPDF) PageHeight() float64 {
	if s.Page.IsLandscape {
		return s.Page.Width
	}
	return s.Page.Height
}

// PageWidth returns the current PDF document width depending on orientation Portrait or Landscape.
// Important note: this is the ENTIRE page and not with margins.
func (s *SimPDF) PageWidth() float64 {
	if s.Page.IsLandscape {
		return s.Page.Height
	}
	return s.Page.Width
}

// SetStyle sets the specific style information for the PDF document to utilize. If fontOnly equals false.
// Then all style options are loaded and not just the font style.
// Loading of font style only allows for quicker calculations of column widths when creating a table.
func (s *SimPDF) SetStyle(style models.Styles, fontOnly bool) {
	//font
	s.PDF.SetTextColor(int(style.Color.Red), int(style.Color.Green), int(style.Color.Blue))
	s.PDF.SetFont(style.Font.Name, style.TextVariant.ToPDF(), style.TextSize)
	if !fontOnly {
		s.PDF.SetLineWidth(style.Border.Width.Bottom)
		s.PDF.SetDrawColor(int(style.Border.Color.Red), int(style.Border.Color.Green), int(style.Border.Color.Blue))
		s.PDF.SetFillColor(int(style.BackgroundColor.Red), int(style.BackgroundColor.Green), int(style.BackgroundColor.Blue))
	}
}

// StringWidth returns the width fo a given string utilizing the current font size loaded.
func (s *SimPDF) StringWidth(text string) float64 {
	return s.PDF.GetStringWidth(text)
}

// CheckBottom ensures that the bottom of the page including the bottom margin is not going to be passed.
// true equals it does pass the bottom of the page.
// false equals it is not going to pass the bottom of the page.
func (s *SimPDF) CheckBottom() bool {
	cY := s.PDF.GetY()
	if (cY + s.Margin.Bottom) > s.PageHeight() {
		return true
	}
	return false
}

func (s *SimPDF) fontReset(style models.Styles) {
	if style.Name == "" {
		sty, err := s.StyleName("Normal")
		internal.IfError("fontReset private", err, false)
		s.SetStyle(sty, true)
	} else {
		s.SetStyle(style, true)
	}
}

// ChangeFont changes the current font to the one included in the style variable. This also
// includes the color of the text as well.
func (s *SimPDF) ChangeFont(style models.Styles) {
	s.PDF.SetTextColor(int(style.Color.Red), int(style.Color.Green), int(style.Color.Blue))
	s.PDF.SetFont(style.Font.Name, style.TextVariant.ToPDF(), style.TextSize)
}

// HeadingStart this function is called before placing a heading into the PDF document.
// Useful for adding a new line, maybe a bookmark, or anything else.
func (s *SimPDF) HeadingStart(styleType, text string) {
	if !strings.Contains(strings.ToLower(styleType), "subtitle") {
		s.AddNewLine(0)
		// s.AddBookmark(styleType, text)
	}
}

// HeadingEnd this function is called after placing a heading into the PDF document.
// Useful for adding a new line, maybe a bookmark, or anything else.
func (s *SimPDF) HeadingEnd(styleType string) {
	if !strings.Contains(strings.ToLower(styleType), "title") {
		s.AddNewLine(0)
	}
}

// AddMargins Adds margins to the current PDF document based upon the margin variable presented.
func (s *SimPDF) AddMargins(margin models.Margins) {
	s.PDF.SetMargins(margin.Left, margin.Top, margin.Right)
}

// WriteCenter centers the current position of X,Y coordinates in the PDF document and then writes
// out the contents of text. align variable is "L" left, "C" center, or "R" right text alignment.
func (s *SimPDF) WriteCenter(styleType string, align models.Alignments, text string) {
	y := s.PageWidth() / 2
	style, _ := s.StyleName(styleType)
	// internal.IfError("WriteCenter public "+err.Error(), err, false)
	s.PDF.SetY(y - (style.LineSize / 2))
	s.Write(styleType, align, text)
}

// Write Simply writes the contents of the variable text to the PDF document as perscribed by the
// styleType variable. align variable is "L" left, "C" center, or "R" right text alignment.
func (s *SimPDF) Write(styleType string, align models.Alignments, text string) {
	if s.CheckBottom() {
		s.AddPageBreak()
	}
	if !strings.Contains(strings.ToLower(styleType), "normal") {
		s.HeadingStart(styleType, text)
	}
	style, _ := s.StyleName(styleType)
	clean := s.Parser(styleType, align, text)
	if len(clean) > 0 {
		style, _ := s.StyleName(styleType)
		s.ChangeFont(style)
		s.PDF.WriteAligned(0, style.LineSize, clean, align.ToPDF())

	}
	if !strings.Contains(strings.ToLower(styleType), "normal") {
		s.HeadingEnd(styleType)
	}
	reset, _ := s.StyleName("Normal")
	s.AddBottomLine(style)
	s.fontReset(reset)
}

// ChangePage changes the page size to the one set as page variable must be a models.Pages
// object for ease of use and will use the Design units pt. Page dimensions should be PORTRAIT
// by default. This is dictated by gofpdf.PDF .
func (s *SimPDF) ChangePage(page models.Pages) {
	s.Page = page
	s.PDF.AddPageFormat(page.ToPDFOrientation(), gofpdf.SizeType{Wd: page.Width, Ht: page.Height})
}

func (s *SimPDF) init(fontDirectory string) {
	// s.PDF = gofpdf.New(s.Page.ToPDFOrientation(), "pt", s.Page.Size, "")
	s.PDF = gofpdf.NewCustom(&gofpdf.InitType{
		UnitStr:        "pt",
		Size:           gofpdf.SizeType{Wd: s.Page.Width, Ht: s.Page.Height},
		FontDirStr:     fontDirectory,
		OrientationStr: s.Page.ToPDFOrientation(),
	})
	s.PDF.AddPage()
}

// Start The main entry function for creating a PDF document.
// pageType - string sizes
// isLandscape - true means the page orientation is landscape, false is portrait
func (s *SimPDF) Start(pageType string, isLandscape bool, styles []models.Styles, margin models.Margins, customFontDirectory string) {
	s.SetPage(pageType, isLandscape)
	s.Style = styles
	s.Margin = margin
	s.init(customFontDirectory)
	s.AddMargins(s.Margin)
}

// Finish ends the creation of the PDF document and saves it to a file as prescribed in fileOutput.
// If a file already exists at fileOutput it will be moved and appened with ".bak" file extention.
func (s *SimPDF) Finish(fileOutput string) {
	if internal.ValidateFilePath(fileOutput) {
		internal.MoveFilePath(fileOutput, fileOutput+".bak")
	}
	s.PDF.OutputFileAndClose(fileOutput)
}
