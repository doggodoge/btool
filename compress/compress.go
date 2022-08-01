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

func appendExtension(paths []string, extension string) []string {
	var ext []string
	for _, p := range paths {
		ext = append(ext, fmt.Sprintf("%s%s", p, extension))
	}
	return ext
}

func FilesToTar(archiveName string, paths []string) error {
	tarPaths := appendExtension(paths, ".tar.br")
	cmdOptions := []string{"-c", "-f", archiveName}
	cmdOptions = append(cmdOptions, tarPaths...)

	cmd := exec.Command("tar", cmdOptions...)
	if err := cmd.Run(); err != nil {
		fmt.Println("failed to run create tar command")
		return err
	}

	for _, p := range tarPaths {
		err := os.Remove(p)
		if err != nil {
			fmt.Printf("failed to remove file %s\n", p)
			return err
		}
		fmt.Printf("removed %s\n", p)
	}
	
	return nil
}
