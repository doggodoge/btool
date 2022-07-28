package compress

import (
	"fmt"
	"os"
	"os/exec"
)

func tar(path string) error {
	cmd := exec.Command("tar", "-c", "-f", fmt.Sprintf("%s.tar", path), path)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func brotli(path string) error {
	cmd := exec.Command("brotli", path, "-o", fmt.Sprintf("%s.br", path))
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func File(path string) error {
	err := tar(path)
	if err != nil {
		return err
	}
	err = brotli(fmt.Sprintf("%s.tar", path))
	if err != nil {
		return err
	}
	err = os.Remove(fmt.Sprintf("%s.tar", path))
	if err != nil {
		return err
	}
	return nil
}
