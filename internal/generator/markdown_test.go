package generator

import (
	"strings"
	"testing"
)

func TestMarkdownToHTML(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Basic bold text",
			input:    "This is **bold**",
			expected: "<strong>bold</strong>",
		},
		{
			name:  "XSS Sanitization check",
			input: "Hello <script>alert('hacked');</script> World",
			// The bluemonday sanitizer should remove the script tags. 
			// In UGCPolicy, the script content might remain as text but it's sanitized.
			expected: "Hello alert(&#39;hacked&#39;); World",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := MarkdownToHTML(tt.input)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			// Convert to string and trim whitespace for easy comparison
			outStr := strings.TrimSpace(string(output))

			if !strings.Contains(outStr, tt.expected) {
				t.Errorf("Expected output to contain %q, but got %q", tt.expected, outStr)
			}
		})
	}
}
