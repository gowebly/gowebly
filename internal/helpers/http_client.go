package helpers

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gowebly/gowebly/internal/constants"
)

type Download struct {
	URL, OutputFile string
}

// DownloadFiles downloads and saves a file with name from the given URL.
func DownloadFiles(files []Download) error {
	// Create a new HTTP client with options.
	// For more information, see https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts/
	client := &http.Client{
		Timeout: 15 * time.Second,
		Transport: &http.Transport{
			TLSHandshakeTimeout:   10 * time.Second,
			ResponseHeaderTimeout: 10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}

	// Create a new context and a cancel function.
	ctx, cancel := context.WithCancel(context.Background())

	// Create error channel.
	errChan := make(chan error, len(files))

	for _, f := range files {
		// Run download process in goroutine.
		go func(f Download, client *http.Client) {
			// Notify the main goroutine about the completion of the current goroutine.
			defer func() { errChan <- nil }()

			// Create a new HTTP request with context.
			req, err := http.NewRequestWithContext(ctx, http.MethodGet, f.URL, http.NoBody)
			if err != nil {
				errChan <- err
				return
			}

			// Download file from the given URL.
			resp, err := client.Do(req)
			if err != nil {
				errChan <- err
				return
			}

			// Check server response.
			if resp.StatusCode != http.StatusOK {
				errChan <- fmt.Errorf(constants.ErrorHTTPDownloadFile, f.URL, resp.StatusCode)
				return
			}

			// Read response body.
			data, err := io.ReadAll(resp.Body)
			if err != nil {
				errChan <- err
				return
			}

			// Create a temp file for the downloaded data.
			if err := MakeFile(File{f.OutputFile, data}); err != nil {
				errChan <- err
				return
			}

			// Close response body.
			if err := resp.Body.Close(); err != nil {
				errChan <- err
				return
			}
		}(f, client)
	}

	// Wait for all goroutines to complete.
	for range files {
		// Cancel the context if an error occurred.
		if err := <-errChan; err != nil {
			cancel()
			return err
		}
	}

	// Cancel the context if the process is successes.
	cancel()

	return nil
}
