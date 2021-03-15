package main

import (
	"github.com/braddschick/simpdf"
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
	pdf.WriteCenter("Title", models.Alignments{Center: true}, "This is my Title")
	// pdf.Write writes text in the current X, Y position of the document
	// This is identical to pdf.WriteCenter for arguments
	pdf.Write("Subtitle", *l, "subtitles are great for dates 2020 March 02")
	// pdf.AddPageBreak creates a new Page Break for the document
	pdf.AddPageBreak()
	pdf.Write("Heading 1", *l, "Important H1")
	pdf.Write("Normal", *l, "Summary usually goes here for the description under the H1. This is really the best I can come up with to fill the paragraph. Well I am almost done but thought a little more to ensure it goes through the breaks.")
	pdf.Write("Heading 3", *l, "Sub H3")
	pdf.Write("Normal", *l, "Summary usually goes here for the description under the H1. This is really the best I can come up with to fill the paragraph. Well I am almost done but thought a little more to ensure it goes through the breaks.")
	pdf.AddPageBreak()
	pdf.Write("Heading 2", *l, "Not as important as the H1")
	pdf.Write("Normal", *l, "Summary usually goes here for the description under the H1. This is really the best I can come up with to fill the paragraph. Well I am almost done but thought a little more to ensure it goes through the breaks.")
	pdf.Write("Normal", *l, "Summary usually goes here for the description under the H1. This is really the best I can come up with to fill the paragraph. Well I am almost done but thought a little more to ensure it goes through the breaks.")
	pdf.Write("Heading 3", *l, "Important H3")
	// Demonstrates the use of inline Bold (__text__), Underline (_#text#_), and Italics (_*text*_)
	pdf.Write("Normal", *l, "__Summary__ usually goes _#here#_ for the description under the _*H1*_. This is really the best _#I#_ can come up with to fill the paragraph. Well I am almost done but thought a little __more__ to ensure it goes through the _#breaks#_.")
	// pdf.Finish creates the PDF document to the file path listed "./test.pdf"
	pdf.Finish("./test.pdf")
}
