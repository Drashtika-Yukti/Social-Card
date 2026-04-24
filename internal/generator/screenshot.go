package generator

import (
	"context"
	"os"
	"path/filepath"

	"github.com/chromedp/chromedp"
)

// Screenshot HTML content using headless chrome
func Screenshot(htmlContent string) ([]byte, error) {
	tmpFile, err := os.CreateTemp("", "social-forge-*.html")
	if err != nil {
		return nil, err
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.WriteString(htmlContent); err != nil {
		return nil, err
	}
	tmpFile.Close()

	fileUrl := "file:///" + filepath.ToSlash(tmpFile.Name())

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var buf []byte
	err = chromedp.Run(ctx,
		chromedp.Navigate(fileUrl),
		chromedp.EmulateViewport(800, 420),
		chromedp.FullScreenshot(&buf, 100),
	)

	if err != nil {
		return nil, err
	}
	return buf, nil
}
