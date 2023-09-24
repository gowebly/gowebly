package helpers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
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

	for _, f := range files {
		// Download and save file from URL.
		if err := func(f Download, client *http.Client) error {
			// Download file from the given URL.
			resp, err := client.Get(f.URL)
			if err != nil {
				return err
			}
			defer resp.Body.Close()

			// Check server response.
			if resp.StatusCode != http.StatusOK {
				return fmt.Errorf("http: can't download file from URL '%s' (code %d)", f.URL, resp.StatusCode)
			}

			// Create a temp file for download data.
			file, err := os.Create(f.OutputFile)
			defer file.Close()

			// Rename downloaded file.
			_, err = io.Copy(file, resp.Body)
			if err != nil {
				return err
			}

			return nil
		}(f, client); err != nil {
			return err
		}
	}

	return nil
}
