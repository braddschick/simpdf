package simpdf

import (
	"errors"
	"fmt"
	"strings"

	"github.com/braddschick/simpdf/internal"
	"github.com/jung-kurt/gofpdf"
)

// Additions is used for any additions to a document such as Images
type Additions interface {
	New(filePath string, width, height float64) error
	Validate() bool
}

// Images struct for Simplifing the addition of an image to the PDF document.
type Images struct {
	// FilePath string of the image file
	FilePath string
	// Extension image file extension. *This is automatically filled
	Extension string
	// Width of the image to be placed in the document
	Width float64
	// Height of the image to be placed in the document
	Height float64
}

// Validate will ensure the image is accessible in the file system
func (i *Images) Validate() bool {
	if internal.ValidateFilePath(i.FilePath) {
		return true
	}
	internal.IfError(fmt.Sprintf("Failed to find image file path (%s) was NOT found", i.FilePath), nil, false)
	return false
}

// New is a function of Images and allows for the automatic validation of the Image file
func (i *Images) New(filePath string, width, height float64) error {
	i.FilePath = filePath
	i.Width = width
	i.Height = height
	i.Extension = internal.FileExtension(filePath)
	if !i.Validate() {
		return errors.New("failed to find the image or was not allowed to access image file")
	}
	return nil
}

// AddImageXY Simply allows the adding of an image to the specifc X Y coordinates
func (s *SimPDF) AddImageXY(image Images, x, y float64) {
	if image.Extension == "PNG" {
		s.PDF.ImageOptions(image.FilePath, x, y, image.Width, image.Height, false, gofpdf.ImageOptions{ImageType: image.Extension, ReadDpi: true}, 0, "")
	} else {
		s.PDF.Image(image.FilePath, x, y, image.Width, image.Height, false, "", 0, "")
	}
}

// AddImageCurrent Adds the image to the current X Y coordinates with a padding of 5 pts.
// Error checking is in place to add a new line to the document if the image runs past the boundaries
// of the document. If a new line has been placed any padding has been removed.
func (s *SimPDF) AddImageCurrent(image Images) {
	x, y := s.PDF.GetXY()
	width := s.PageWidth()
	if (s.Margin.Left+s.Margin.Right)+(x+image.Width) > width {
		s.AddNewLine()
		x, y = s.PDF.GetXY()
	} else {
		x = x + 5
		y = y + 5
	}
	s.AddImageXY(image, x, y)
}

// AddImageStandardPosition Adds an image to the Standard Position "tr, tc, tl, ...".
// If the Standard Position contains "r" - Right the Image Width is minus'ed from the X.
// If the Standard Position contains "b" - Bottom the Image Height is minus'ed from the Y.
func (s *SimPDF) AddImageStandardPosition(image Images, stdPosition string) {
	x, y := s.StandardPosition(stdPosition)
	if strings.Contains(strings.ToLower(stdPosition), "r") {
		x = x - image.Width
	}
	if strings.Contains(strings.ToLower(stdPosition), "b") {
		y = y - image.Height
	}
	s.AddImageXY(image, x, y)
}
