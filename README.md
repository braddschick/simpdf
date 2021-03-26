![SimPDF simple PDF development](./test/images/simpdf.png)

# SimPDF

Simplify gofpdf.Fpdf document generation with with this wrapper for use in golang.

## Simplify and Expand

SimPDF has a simple goal of simplifying the [gofpdf.Fpdf](https://github.com/jung-kurt/gofpdf/) package for use in golang
for the masses. The gofpdf package is a great PDF creator and has many various options that are hard to utilize for novice
programmers. This hopes to be solved, _slightly_, by this package.

## Features
- Table handling
    - Alternate Row Styling
    - Header Row Styling
    - Column Width
        - Fixed
        - Distribute Evenly the width of the page
        - Auto
- Images
    - Change Width/Height proportionately
    - Position
        - Inset with Text
        - Standard Positioning
- Styling
    - Sets of styles for easy switching of styles
    - WEB CSS color values *HEX* to RGB Color
    - Complete Google Material Colors available
- Paper Sizes Available
    - A1 - A5
    - Letter, Legal, Tabloid, and Ledger
    - ANSI A
- Orientation switching in document
    - Allows for multiple pages to be a different orientation than the original starting orientation

## Installation

``` bash

go get "github.com/braddschick/simpdf"

```

## Examples

``` go

// pdf.Start - Starts the creation of a PDF document
var pdf simpdf.SimPDF
// "Letter" is the page size could also be Legal, Tabloid, Ledger, A1 - A5
// false means this will be in portrait orientation, true means landscape
// defaults.BasicStyle contains the styles to be used in the document you can create your own
// defaults.Narrow_Margins is the margin information for the document
// "" if you want to use custom fonts this will be the custom font directory path
pdf.Start("Letter", false, defaults.BasicStyle, defaults.Narrow_Margins, "")
pdf.Details("Title", "Author", "Subject", "More Keywords, here, here")
l := &models.Alignments{Left: true}
// pdf.WriteCenter writes in the center of the document great for title pages
// "Title" denotes the style of text to be used. This is the name of the models.Style to utilize
// models.Alignments{Center: true} will center the text to the page, also can be Left: true or Right: true
// "This is my Title" is the text to be written out to the document
pdf.WriteCenter("Title", *l, "Title \"centered to the page\"")
// Manual page break
pdf.AddPageBreak()
pdf.Write("Normal", *l, "Adding an image to a \"Standard Position\" is easy as well. Top Left, _#tl#_, or Top Center, _#tc#_, or Top Right, _#tr#_, and is also available in Center or Bottom variations.")
pdf.Finish("./simple_example.pdf")

```

### Complete Examples

See the [Examples](./examples) directory for various examples of how to use SimPDF. 

- [Simple Example](./examples/simple.go) - Basic operations Write/Center, Alignment, Instantiation, and completion
