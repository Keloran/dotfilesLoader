package dots

import (
	"fmt"

	"github.com/Keloran/dotfilesLoader/files"
)

type Dots struct {
	Username   string
	Github     files.Github
	Force      bool
	Skip       bool
	CurrentDir string
}

func (d Dots) Install() error {
  type dots struct {
    Location string
    Name string
  }

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

			dd := []dots{}
			prefixLen := len(zipedFiles[0])
			for _, file := range zipedFiles[1:] {
			  if file[(prefixLen + 1):(prefixLen + 2)] == "." {
			    ddd := dots{
			      Location: file,
			      Name: file[(prefixLen + 1):],
          }

			    dd = append(dd, ddd)
        }
      }

      for _, dot := range dd {
        err = files.Copy(dot.Location, fmt.Sprintf("%s/%s", d.CurrentDir, dot.Name))
        if err != nil {
          return fmt.Errorf("Dotfiles copy: %w", err)
        }
      }
		}

		err = files.Cleanup("/tmp/github")
		if err != nil {
      return fmt.Errorf("dotfiles installed cleanup: %w", err)
    }
	}

	fmt.Print("dotfiles installed\n")
	return nil
}
