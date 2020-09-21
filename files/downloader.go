package files

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type Github struct {
	Username      string
	Repository    string
	GivenLocation string
}

type Downloader struct {
	GithubDetails Github
	Skip          bool

	StoredLocation string
}

func (d Downloader) Github() (string, error) {
	if d.GithubDetails.Username == "" {
		return "", fmt.Errorf("Downloader Github Username blank")
	}

	if d.GithubDetails.Repository == "" {
		return "", fmt.Errorf("Downloader Github Repository blank")
	}

	if !d.Skip {
		d, err := d.getFiles(fmt.Sprintf("https://github.com/%s/%s/archive/master.zip", d.GithubDetails.Username, d.GithubDetails.Repository), "github.zip")
		return d.StoredLocation, err
	}

	if d.GithubDetails.GivenLocation != "" {
		return d.GithubDetails.GivenLocation, nil
	}

	return "/tmp/github", nil
}

func (d Downloader) getFiles(location, filename string) (Downloader, error) {
	fileLoc := fmt.Sprintf("/tmp/%s", filename)

	resp, err := http.Get(location)
	if err != nil {
		return d, fmt.Errorf("Downloader getFiles: %w", err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Printf("Downloader getFiles close network %v", err)
		}
	}()

	out, err := os.Create(fileLoc)
	if err != nil {
		return d, fmt.Errorf("Downloader getFiles create: %w", err)
	}
	defer func() {
		if err := out.Close(); err != nil {
			fmt.Printf("Downloader getFiles close file: %v", err)
		}
	}()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return d, fmt.Errorf("Downloader getFiles copy: %w", err)
	}
	d.StoredLocation = fileLoc

	return d, nil
}
