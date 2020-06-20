package dots

import (
  "fmt"
)

type Dots struct {
  Username string
  Github bool
  Force bool
}

func (d Dots) Install() error {
  return fmt.Errorf("tester")

  return nil
}
