package htmlformat

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

type FormatterConfig struct {
	Indent string
}

type Formatter struct {
	config FormatterConfig
}

var bNewLine = []byte("\n")

func NewFormatter(config FormatterConfig) *Formatter {
	if config.Indent == "" {
		config.Indent = "  "
	}

	return &Formatter{config: config}
}

func (f *Formatter) FormatString(raw string) (pretty string, e error) {
	r := strings.NewReader(raw)
	w := &strings.Builder{}
	e = f.Format(r, w)

	return w.String(), e
}

func (f *Formatter) Format(input io.Reader, output io.Writer) error {
	z := html.NewTokenizer(input)
	depth := 0
	start := true
	prevToken := html.CommentToken
	bIndent := []byte(f.config.Indent)

	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			break
		}

		tokenString := string(z.Raw())

		// strip away newlines
		if tt == html.TextToken {
			stripped := strings.Trim(tokenString, "\n")
			if len(stripped) == 0 {
				continue
			}
		}

		if tt == html.EndTagToken {
			depth -= 1
		}

		if tt != html.TextToken {
			if !start && prevToken != html.TextToken {
				output.Write(bNewLine)
				for i := 0; i < depth; i++ {
					output.Write(bIndent)
				}
			}
		}

		output.Write([]byte(tokenString))

		// last token
		if tt == html.StartTagToken {
			depth += 1
		}
		prevToken = tt
		start = false
	}

	return nil
}
