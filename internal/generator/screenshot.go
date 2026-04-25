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

	// Configure Chromedp for Linux/Docker environments (No Sandbox is required)
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.NoSandbox,
		chromedp.DisableGPU,
		chromedp.Flag("disable-setuid-sandbox", true),
		chromedp.Flag("disable-dev-shm-usage", true),
	)

	allocCtx, cancelAlloc := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancelAlloc()

	ctx, cancel := chromedp.NewContext(allocCtx)
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
