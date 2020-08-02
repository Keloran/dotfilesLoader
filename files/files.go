package files

import (
	"fmt"
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
	}

	if err = os.Remove(input); err != nil {
		return fmt.Errorf("Files cleanup remove: %w", err)
	}

	return nil
}
