package helpers

import (
	"fmt"
	"io"
	"net/http"
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

			// Read response body.
			data, err := io.ReadAll(resp.Body)
			if err != nil {
				return err
			}

			// Create a temp file for download data.
			if err := MakeFiles(
				[]File{
					{
						f.OutputFile, data,
					},
				},
			); err != nil {
				return err
			}

			return nil
		}(f, client); err != nil {
			return err
		}
	}

	return nil
}
