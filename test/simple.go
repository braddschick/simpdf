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
	// defaults.BasicStyle[5].Font.Name = "courier"
	pdf.Start("Letter", false, defaults.BasicStyle, defaults.Narrow_Margins, "")
	pdf.Details("Testing Functions", "Author", "Testing Subject", "More Keywords, here, here")
	simpdf.Header.LeftContent.Text = pdf.Title
	simpdf.Header.LeftContent.Style = defaults.Basic_Table
	bold := defaults.Basic_Table
	bold.TextVariant.Bold = true
	simpdf.Header.RightContent.Style = bold
	simpdf.Header.RightContent.PageNumber = true
	simpdf.Header.RightContent.Text = "|"
	pdf.SetHeader()
	simpdf.Footer.CenterContent.Style = defaults.Basic_Table
	simpdf.Footer.CenterContent.Text = "CONFIDENTIAL"
	pdf.SetFooter()
	l := &models.Alignments{Left: true}
	// pdf.WriteCenter writes in the center of the document great for title pages
	// "Title" denotes the style of text to be used. This is the name of the models.Style to utilize
	// models.Alignments{Center: true} will center the text to the page, also can be Left: true or Right: true
	// "This is my Title" is the text to be written out to the document
	pdf.WriteCenter("Title", *l, "Title \"centered to the page\"")
	// pdf.Write writes text in the current X, Y position of the document
	// This is identical to pdf.WriteCenter for arguments
	pdf.Write("Subtitle", *l, "subtitle style")
	// pdf.AddPageBreak creates a new Page Break for the document
	pdf.AddPageBreak()
	pdf.Write("Heading 1", *l, "Heading One")
	pdf.Write("Normal", *l, "This is normal text and again can be styled very easily. This has a left alignment added to it. Also, remember the _*gofpdf.Pdf*_ is always accesible through _*SimPDF.PDF*_.")
	pdf.AddNewLine(0)
	pdf.Write("", models.Alignments{}, "This has no aligment noted and will appear as the default \"left\" alignment. This also has no models.Style.Name provided and will utilize the \"Normal\" style located in the styles provided earlier.")
	pdf.Write("Heading 2", *l, "Heading Level Two")
	pdf.Write("Heading 3", *l, "Heading Level Three")
	goImage, _ := simpdf.NewImage("./images/golang.png", 355, 486)
	pdf.AddImageStandardPosition(goImage, "bc")
	// Manual page break
	pdf.AddPageBreak()
	// Demonstrates the use of inline Bold (__text__), Underline (_#text#_), and Italics (_*text*_)
	pdf.Write("Normal", *l, "Here is formatted text. __Bolded text here__ then we have _#underlined text#_ but you also need to have _*italic text as well*_. This makes it very easy to use text vairants within in texts.")
	// Tables can have alternating row styles
	altRow := defaults.Basic_Table
	altRow.BackgroundColor = colors.Grey
	// Tables also have header row styles that can be applied
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
	pdf.Write("", *l, "Also, tables do not have to have header rows if they are not needed. Just add empty strings to the Tables.Headers string list to ensure the column count is the same.")
	pdf.Write("", *l, "* Note the first column header __CAN__ be blank if required by your table.")
	table.Headers = []string{"", "", ""}
	pdf.AddTable(table, altRow, 0)
	pdf.Write("Normal", *l, "You can add images and also have the size of image modified while constraining the proprotions correctly if required. This image has been placed at the current X, Y position.")
	goImage.ChangeHeight(150)
	pdf.AddImageCurrent(goImage)
	// Adding the Images.Height to simpdf.AddNewLine() ensures there is no text being added on top of the image since I placed it at the current X, Y position.
	pdf.AddNewLine(goImage.Height)
	pdf.Write("Normal", *l, "Adding an image to a \"Standard Position\" is easy as well. Top Left, _#tl#_, or Top Center, _#tc#_, or Top Right, _#tr#_, and is also available in Center or Bottom variations.")
	goImage.ChangeWidth(75)
	pdf.WriteImageInset("Normal", *l, "This needs to have margin left away from the image, but also look decent. _#However#_, I need to ensure it continues with a line break which is why this is so long.", "tl", goImage)
	pdf.Write("normal", *l, "These two images of the GOpher are the same image. Reuse is easy and _*Images.ChangeHeight()*_ or _*Images.ChangeWidth()*_ can easily change the image as needed.")
	// pdf.Finish creates the PDF document to the file path listed "./test.pdf"
	pdf.Finish("./simple_example.pdf")
}
