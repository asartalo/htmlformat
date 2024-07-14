package htmlformat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatterBasic(t *testing.T) {
	html := `<html><head></head><body><a href="http://test.com">Test link</a><p><br/></p></body></html>`
	expected := `<html>
  <head>
  </head>
  <body>
    <a href="http://test.com">Test link</a>
    <p>
      <br/>
    </p>
  </body>
</html>`
	config := FormatterConfig{}
	formatter := NewFormatter(config)
	pretty, _ := formatter.FormatString(html)
	assert.Equal(t, expected, pretty)
}

func TestFormatterWithNewlinesAndCustomIndentation(t *testing.T) {
	html := `<!doctype html><html><head>
<title>Website Title</title>
</head><body>
<div class="random-class">
<h1>I like pie</h1><p>It's true!</p></div>
</body></html>`
	expected := `<!doctype html>
<html>
    <head>
        <title>Website Title</title>
    </head>
    <body>
        <div class="random-class">
            <h1>I like pie</h1>
            <p>It's true!</p>
        </div>
    </body>
</html>`

	config := FormatterConfig{
		Indent: "    ",
	}
	formatter := NewFormatter(config)
	pretty, _ := formatter.FormatString(html)

	assert.Equal(t, expected, pretty)
}
