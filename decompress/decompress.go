package decompress

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func untar(path string) error {
	cmd := exec.Command("tar", "-x", "-f", path)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func unbrotli(path string) error {
	cmd := exec.Command("brotli", "-d", path)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func File(path string) error {
	strippedPath := strings.Split(path, ".")[0]
	err := unbrotli(fmt.Sprintf("%s.tar.br", strippedPath))
	if err != nil {
		return err
	}
	err = untar(fmt.Sprintf("%s.tar", strippedPath))
	if err != nil {
		return err
	}
	err = os.Remove(fmt.Sprintf("%s.tar", strippedPath))
	if err != nil {
		return err
	}
	return nil
}
