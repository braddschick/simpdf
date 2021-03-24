package simpdf

import (
	"regexp"
	"strings"

	"github.com/braddschick/simpdf/pkg/models"
)

var (
	// BoldExp checks to see if the double underscore was used to denote bold for text
	BoldExp = regexp.MustCompile(`(?m)(?:\_\_([\S]+[^\_\_]*)\_\_)`)
	// BoldSub is the html for bolding within group matching
	BoldSub = "<b>$1</b>"
	// StrikeExp is denoted by ~~strike through text~~
	// StrikeExp = regexp.MustCompile(`(?m)(?:\~\~([\S]+[^\~]*)\~\~))`)
	// ItalicExp is denoted by _*italic text*_
	ItalicExp = regexp.MustCompile(`(?m)(?:\_\*([\S]+[^\_\*]*)\*\_)`)
	// ItalicSub is the html for italics within group matching
	ItalicSub = "<i>$1</i>"
	// UnderlineExp is denoted by _#underlined text#_
	UnderlineExp = regexp.MustCompile(`(?m)(?:\_\#([\S]+[^\_\#]*)\#\_)`)
	// UnderlineSub is the html for underlinging within group matching
	UnderlineSub = "<u>$1</u>"
)

type Expressions interface {
	IsMatch(text string) bool
	Parse(text string) (string, bool)
	ReplaceAll(text string) string
}

// ParseGroup holds the regular expression group for substitution purposes
type ParseGroup struct {
	Expression   regexp.Regexp
	Substitution string
}

// IsMatch checks if the simpdf.ParseGroup matches the variable text.
func (p *ParseGroup) IsMatch(text string) bool {
	return p.Expression.MatchString(text)
}

// ReplaceAll replaces all occurrences of the simpdf.ParseGroup with the
// simpdf.ParseGroup.Substitution string. This must follow the standard
// regex group substitution format "$1".
func (p *ParseGroup) ReplaceAll(text string) string {
	return p.Expression.ReplaceAllString(text, p.Substitution)
}

// Parse will replace all matches with proper susbstitution group. This is used
// for replacing the "markdown-like" structures with the proper HTML code for
// rendering.
func (p *ParseGroup) Parse(text string) (string, bool) {
	if p.IsMatch(text) {
		return p.ReplaceAll(text), true
	}
	return text, false
}

func createParseGroup() []ParseGroup {
	var pGroup []ParseGroup
	pGroup = append(pGroup, ParseGroup{
		Expression:   *BoldExp,
		Substitution: BoldSub,
	})
	pGroup = append(pGroup, ParseGroup{
		Expression:   *ItalicExp,
		Substitution: ItalicSub,
	})
	pGroup = append(pGroup, ParseGroup{
		Expression:   *UnderlineExp,
		Substitution: UnderlineSub,
	})
	return pGroup
}

// Parser this checks if there is any Markdown like style requirements in the
// text to be replaced by HTML that gofpdf.PDF understands.
//
// __text__ => Bold text
// _*text*_ => Italic text
// _#text#_ => Underline text
//
// Return string of the text orginally assigned if it does not contain any MD,
// string will be empty if it did contain any MD to be transformed.
func (s *SimPDF) Parser(style string, align models.Alignments, text string) string {
	pGroup := createParseGroup()
	out := text
	needHTML := false
	test := false
	for _, p := range pGroup {
		out, test = p.Parse(out)
		if test {
			needHTML = true
		}
	}
	if needHTML {
		html := s.PDF.HTMLBasicNew()
		sty, _ := s.StyleName(style)
		s.SetStyle(sty, false)
		html.Write(sty.LineSize, strings.Replace(align.ToHTML(), "$1", out, -1))
		return ""
	}
	return text
}
