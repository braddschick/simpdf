package models

import (
	"strconv"
	"strings"
)

type Contents struct {
	PageNumber bool
	Align      Alignments
	Prepend    bool
	Text       string
	Style      Styles
}

func (c *Contents) Empty() bool {
	return strings.Trim(c.Text, " ") == ""
}

func (c *Contents) AddPage(pg int) string {
	if c.PageNumber {
		if c.Prepend {
			return strconv.Itoa(pg) + " " + c.Text
		}
		return c.Text + " " + strconv.Itoa(pg)
	}
	return c.Text
}

func (c *Contents) ToHTML(pg int) string {
	return strings.Replace(c.Align.ToHTML(), "$1", c.AddPage(pg), -1)
}
