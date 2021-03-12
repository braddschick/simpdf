package simpdf

import (
	"fmt"
	"math"
	"strings"

	"github.com/braddschick/simpdf/internal"
	"github.com/braddschick/simpdf/pkg/models"
)

// Tables struct is a simple object for inclusion of a Table into the PDF document.
type Tables struct {
	// Headers simple striing list of the headers with the alignment**content format
	// Example: C**Titles will result as Titles with Center text alignment for the Column 1
	Headers []string
	// HeaderStyle contains the models.Styles that will depict the Header Row of the table
	HeaderStyle models.Styles
	// Rows contains each row of the table. The alignment**content format is observed.
	// Each string list is a row {"L**Column1", "C**Column2", "R**Column3"}
	Rows [][]string
	// RowStyle contains the models.Styles that will depict each data row of the table.
	RowStyle models.Styles
	// HasAlternating denotes if there was an alternating style for the data rows.
	HasAlternating bool
	// AlternatingRowStyle contains the models.Styles that will depict each even data row.
	AlternatingRowStyle models.Styles
	// MaxColWidth float64 list that has the Maximum Column Width of each column based on data in the table.
	MaxColWidth []float64
}

// BreakTableAlignment Tables function for spliting the alignment from the cell text.
// C, L, R - Center, Left, Right is alignment of the cell contents.
// Alignment must precede the cell contents followed by "**"
// Example of this is "C**Cell text goes here". This means the cell text will be Centered.
func BreakTableAlignment(str string) (string, string) {
	strs := strings.Split(str, "**")
	if len(strs) < 2 {
		strs = []string{"L", str}
	}
	return strings.ToUpper(strs[0]), strs[1]
}

// TableColumnWidth will determine the width of each column at the max width of the contents
// plus 6 pts of padding.
// This should NOT be used directly but is provided for context. Use AddTable() instead.
func (s *SimPDF) TableColumnWidth(table Tables) []float64 {
	iCols := make([]float64, len(table.Headers))
	s.SetStyle(table.HeaderStyle, true)
	for i := range iCols {
		iCols[i] = math.Round(s.StringWidth(table.Headers[i])) + 6
	}
	s.SetStyle(table.RowStyle, true)
	for _, str := range table.Rows {
		for ix, j := range iCols {
			vW := math.Round(s.StringWidth(str[ix])) + 6
			if vW > j {
				iCols[ix] = vW
			}
		}
	}
	return iCols
}

// AddTableHeader Adds the table header row to the PDF document. If fixWidth is not 0 then
// all cells will be set to the fixed width of the fixWidth value.
// This should NOT be used directly but is provided for context. Use AddTable() instead.
func (s *SimPDF) AddTableHeader(table Tables, fixWidth float64) {
	s.SetStyle(table.HeaderStyle, false)
	for i, r := range table.Headers {
		b := fmt.Sprintf("%g", table.HeaderStyle.Border.Width.Top)
		var w float64
		if fixWidth == 0 {
			w = table.MaxColWidth[i]
		} else {
			w = fixWidth
		}
		align, str := BreakTableAlignment(r)
		s.PDF.CellFormat(w, table.HeaderStyle.LineSize, str, b, 0, align, true, 0, "")
	}
	s.AddNewLine()
}

// AddTableRows Adds the table rows to the PDF document. If fixWidth is not 0 then
// all cells will be set to the fixed width of the fixWidth value.
// This should NOT be used directly but is provided for context. Use AddTable() instead.
func (s *SimPDF) AddTableRows(table Tables, fixWidth float64) {
	for ir, r := range table.Rows {
		b := fmt.Sprintf("%g", table.RowStyle.Border.Width.Top)
		if ir%2 == 1 {
			if table.HasAlternating {
				s.SetStyle(table.AlternatingRowStyle, false)
			}
		} else {
			s.SetStyle(table.RowStyle, false)
		}
		for ix, n := range r {
			var w float64
			if fixWidth == 0 {
				w = table.MaxColWidth[ix]
			} else {
				w = fixWidth
			}
			align, str := BreakTableAlignment(n)
			s.PDF.CellFormat(w, table.RowStyle.LineSize, str, b, 0, align, true, 0, "")
		}
		s.AddNewLine()
	}
	sty, err := s.StyleName("Normal")
	internal.IfError("AddTableRows secure", err, false)
	s.SetStyle(sty, false)
}

// AddTable Simply adds the table to the PDF document. This is the main function for adding a
// table to the document. If fixWidth is not 0 then all cells will be set to the fixed width of
// the fixWidth value. If it is 0 then the width will be dependent on the cell contents.
func (s *SimPDF) AddTable(table Tables, altRowColor models.Styles, fixedWidth float64) {
	if altRowColor.Name != "" {
		table.HasAlternating = true
		table.AlternatingRowStyle = altRowColor
	}
	s.AddNewLine()
	s.AddNewLine()
	table.MaxColWidth = s.TableColumnWidth(table)
	s.AddTableHeader(table, fixedWidth)
	s.AddTableRows(table, fixedWidth)
	style, err := s.StyleName("Normal")
	internal.IfError("AddTable public", err, false)
	s.fontReset(style)
	s.AddNewLine()
}

// DistributeColumnsEvenly returns a fixed width size that would allow the columns to be evenly
// distributed across the page width of the PDF document.
func (s *SimPDF) DistributeColumnsEvenly(numCols float64) float64 {
	return (s.PageWidth() - (s.Margin.Left + s.Margin.Right)) / numCols
}
