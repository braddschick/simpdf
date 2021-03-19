package main

import (
	"github.com/braddschick/simpdf"
	"github.com/braddschick/simpdf/pkg/colors"
	"github.com/braddschick/simpdf/pkg/defaults"
	"github.com/braddschick/simpdf/pkg/models"
)

func main() {
	var pdf simpdf.SimPDF
	// pdf.Start - Starts the creation of a PDF document
	// "Letter" is the page size could also be Legal, Tabloid, Ledger, A1 - A5
	// false means this will be in portrait orientation, true means landscape
	// defaults.BasicStyle contains the styles to be used in the document you can create your own
	// defaults.Narrow_Margins is the margin information for the document
	// "" if you want to use custom fonts this will be the custom font directory path
	pdf.Start("Letter", false, defaults.BasicStyle, defaults.Narrow_Margins, "")
	l := &models.Alignments{Left: true}
	// pdf.WriteCenter writes in the center of the document great for title pages
	// "Title" denotes the style of text to be used. This is the name of the models.Style to utilize
	// models.Alignments{Center: true} will center the text to the page, also can be Left: true or Right: true
	// "This is my Title" is the text to be written out to the document
	pdf.WriteCenter("Title", models.Alignments{Center: true}, "Title \"centered\"")
	// pdf.Write writes text in the current X, Y position of the document
	// This is identical to pdf.WriteCenter for arguments
	pdf.Write("Subtitle", *l, "subtitle which can be styled differently")
	// pdf.AddPageBreak creates a new Page Break for the document
	pdf.AddPageBreak()
	pdf.Write("Heading 1", *l, "Heading One")
	pdf.Write("Normal", *l, "This is normal text and again can be styled very easily. This has a left alignment added to it. Also, remember the gofpdf.Pdf is always accesible through SimPDF.PDF.")
	pdf.AddNewLine(0)
	pdf.AddNewLine(0)
	pdf.Write("", models.Alignments{}, "This has no aligment noted and will appear as the default \"left\" alignment. This also has no models.Style.Name provided and will utilize the \"Normal\" style located in the styles provided earlier.")
	pdf.Write("Heading 2", *l, "Heading Level Two")
	pdf.Write("Heading 3", *l, "Heading Level Three")
	pdf.AddPageBreak()
	// Demonstrates the use of inline Bold (__text__), Underline (_#text#_), and Italics (_*text*_)
	pdf.Write("Normal", *l, "Here is formatted text. __Bolded text here__ then we have _#underlined text#_ fbut you also need to _*italic text as well*_. This makes it very easy to use text vairants within in texts.")
	// pdf.Finish creates the PDF document to the file path listed "./test.pdf"
	altRow := defaults.Basic_Table
	altRow.BackgroundColor = colors.Grey
	headerRow := defaults.Basic_Table
	headerRow.BackgroundColor = colors.Black
	headerRow.Color = colors.White
	headerRow.Font.Name = "Arial"
	headerRow.TextSize = headerRow.TextSize + 6
	headerRow.LineSize = headerRow.LineSize + 6
	headerRow.TextVariant.Bold = true
	table := simpdf.Tables{
		Headers:     []string{"C**Character", "C**Premiered", "C**Salary"},
		HeaderStyle: headerRow,
		RowStyle:    defaults.Basic_Table,
		Rows: [][]string{
			{"L**Mickey Mouse", "C**1928", "L**$  3,000,000,000"},
			{"L**Popeye", "C**1919", "L**$  500,000"},
			{"L**Donald Duck", "C**1934", "L**$  5,000,000"},
		},
	}
	pdf.Write("", *l, "This is a simple table with _*No fixed width*_.")
	pdf.AddTable(table, altRow, 0)
	pdf.Write("", *l, "This is a simple table with _*Fixed width*_ _#140pts#_.")
	pdf.AddTable(table, altRow, 140)
	pdf.Write("", *l, "This is a simple table with _*Ditstribute Evenly column width*_.")
	pdf.AddTable(table, altRow, pdf.DistributeColumnsEvenly(3))
	pdf.AddPageBreak()
	pdf.Write("Normal", *l, "You can add images and also have the size of image modified while constraining the proprotions correctly if required. This image has been placed at the current X, Y position.")
	goImage, _ := simpdf.NewImage("./images/golang.png", 355, 486)
	goImage.ChangeHeight(200)
	pdf.AddImageCurrent(goImage)
	// Adding the Images.Height to simpdf.AddNewLine() ensures there is no text being added on top of the image since I placed it at the current X, Y position.
	pdf.AddNewLine(goImage.Height)
	pdf.Write("Normal", *l, "Adding an image to a \"Standard Position\" is easy as well. Top Left, _#tl#_, or Top Center, _#tc#_, or Top Right, _#tr#_, and is also available in Center or Bottom variations.")
	pdf.Write("Normal", *l, "The little gopher is located at the \"Bottom Left\" by using _*bl*_.")
	goImage.ChangeWidth(100)
	// pdf.AddImageStandardPosition(goImage, "bl")
	pdf.WriteImageInset("Normal", *l, "This needs to have margin left away from the image, but also look decent. _#However#_, I need to ensure it continues with a line break which is why this is so long.", "tr", goImage)
	pdf.Finish("./simple_example.pdf")
}
