package files

import (
	"fmt"
  "io"
  "os"
)

func Cleanup(input string) error {
	f, err := os.Open(input)
	if err != nil {
		return fmt.Errorf("Files cleanup open: %w", err)
	}
	defer func() {
		err := f.Close()
		if err != nil {
			fmt.Printf("Files cleanup close: %v", err)
		}
	}()

	info, err := f.Stat()
	if err != nil {
		return fmt.Errorf("Files cleanup stat: %w", err)
	}
	if info.IsDir() {
		err = os.RemoveAll(input)
		if err != nil {
			return fmt.Errorf("Files cleanup removeAll: %w", err)
		}
		return nil
	}

	if err = os.Remove(input); err != nil {
		return fmt.Errorf("Files cleanup remove: %w", err)
	}

	return nil
}

func Copy(input, location string) error {
  stat, err := os.Stat(input)
  if err != nil {
    return fmt.Errorf("Files copy stat: %w", err)
  }
  if !stat.Mode().IsRegular() {
    return fmt.Errorf("Files copy not a file: %s", input)
  }

  src, err := os.Open(input)
  if err != nil {
    return fmt.Errorf("Files copy open: %w", err)
  }
  defer func() {
    err := src.Close()
    if err != nil {
      fmt.Printf("Files copy close input: %v", err)
    }
  }()

  dest, err := os.Create(location)
  if err != nil {
    return fmt.Errorf("Files copy create: %w", err)
  }
  defer func() {
    err := dest.Close()
    if err != nil {
      fmt.Printf("Files copy close dest: %v", err)
    }
  }()
  _, err = io.Copy(dest, src)
  if err != nil {
    return fmt.Errorf("Files copy copy: %w", err)
  }

  return nil
}
