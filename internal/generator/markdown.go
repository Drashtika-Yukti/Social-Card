package generator

import (
	"bytes"
	"html/template"

	"github.com/microcosm-cc/bluemonday"
	"github.com/yuin/goldmark"
)

// MarkdownToHTML converts markdown text to sanitized HTML
func MarkdownToHTML(md string) (template.HTML, error) {
	var buf bytes.Buffer
	if err := goldmark.Convert([]byte(md), &buf); err != nil {
		return "", err
	}

	// Sanitize output to prevent XSS in the headless browser
	p := bluemonday.UGCPolicy()
	sanitized := p.SanitizeBytes(buf.Bytes())

	return template.HTML(sanitized), nil
}
