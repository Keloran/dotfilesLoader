package files

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func Extract(input, output string) ([]string, error) {
	var files []string

	r, err := zip.OpenReader(input)
	if err != nil {
		return files, fmt.Errorf("Extractor extract openReader: %w", err)
	}
	defer func() {
		if err := r.Close(); err != nil {
			fmt.Printf("Extractor extract close zip: %v", err)
		}
	}()

	folderName := fmt.Sprintf("/tmp/%s", output)
	if err = os.Mkdir(folderName, os.ModePerm); err != nil {
		return files, fmt.Errorf("Extractor extract create folder: %w", err)
	}

	for _, f := range r.File {
		fpath := filepath.Join(folderName, f.Name)

		// ZipSlip
		if !strings.HasPrefix(fpath, filepath.Clean(folderName)+string(os.PathSeparator)) {
			return files, fmt.Errorf("Extractor extract, %v has illegal path in it %s", input, fpath)
		}

		files = append(files, fpath)

		if f.FileInfo().IsDir() {
			if err = os.MkdirAll(fpath, os.ModePerm); err != nil {
				return files, fmt.Errorf("Extractor extract folder create: %w", err)
			}
			continue
		}

		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return files, fmt.Errorf("Extractor extract make file %w", err)
		}

		out, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return files, fmt.Errorf("Extractor extract create file %w", err)
		}

		rc, err := f.Open()
		if err != nil {
			return files, fmt.Errorf("Extractor extract open file %+v, err: %w", f, err)
		}

		_, err = io.Copy(out, rc)
		if err != nil {
			return files, fmt.Errorf("Extractor extract copy file %w", err)
		}

		err = out.Close()
		if err != nil {
			return files, fmt.Errorf("Extractor extract close file %w", err)
		}

		err = rc.Close()
		if err != nil {
			return files, fmt.Errorf("Extractor extract close tmp file %w", err)
		}
	}

	if err = Cleanup(input); err != nil {
		return files, fmt.Errorf("Extractor extract cleanup: %w", err)
	}

	return files, nil
}
