package simpdf

import "github.com/braddschick/simpdf/pkg/models"

type HeaderFooters struct {
	LeftContent   models.Contents
	RightContent  models.Contents
	CenterContent models.Contents
}

func (hf *HeaderFooters) Write(s *SimPDF) {
	html := s.PDF.HTMLBasicNew()
	if !hf.LeftContent.Empty() {
		s.SetStyle(hf.LeftContent.Style, false)
		html.Write(hf.LeftContent.Style.LineSize, hf.LeftContent.ToHTML(s.PDF.PageNo()))
	}
	if !hf.CenterContent.Empty() {
		s.SetStyle(hf.CenterContent.Style, false)
		html.Write(hf.CenterContent.Style.LineSize, hf.CenterContent.ToHTML(s.PDF.PageNo()))
	}
	if !hf.RightContent.Empty() {
		s.SetStyle(hf.RightContent.Style, false)
		html.Write(hf.RightContent.Style.LineSize, hf.RightContent.ToHTML(s.PDF.PageNo()))
	}
}
