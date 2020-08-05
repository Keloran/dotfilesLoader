package dots

import (
	"fmt"

	"github.com/Keloran/dotfilesLoader/files"
)

type Dots struct {
	Username string
	Github   files.Github
	Force    bool
	Skip     bool
}

func (d Dots) Install() error {
	if d.Github.Repository != "" {
		filesLocation, err := files.Downloader{
			GithubDetails: d.Github,
			Skip:          d.Skip,
		}.Github()
		if err != nil {
			return fmt.Errorf("dotfiles install: %w", err)
		}

		if filesLocation == "" {
			return fmt.Errorf("dotfiles install, file location blank")
		}

		if !d.Skip {
			zipedFiles, err := files.Extract(filesLocation, "github")
			if err != nil {
				return fmt.Errorf("dotfiles extract: %w", err)
			}
			fmt.Printf("Files: %+v", zipedFiles)
		}

		err = files.Cleanup("github")
		return fmt.Errorf("dotfiles installed cleanup: %w", err)
	}

	fmt.Printf("Dots %+v\n", d)
	return fmt.Errorf("tester")

	// return nil
}
