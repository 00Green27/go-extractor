package extractor

import (
	"fmt"
	"os"
	"os/exec"
)

type Extractor interface {
	Extract(src, dst string) error
}

type extractor struct{}

func New() Extractor {
	return &extractor{}
}

func (e *extractor) Extract(src, dst string) error {
	srcType, err := mimeType(src)
	if err != nil {
		return err
	}

	switch srcType {
	case "application/x-rar-compressed":
		err = extractRar(src, dst)
	case "application/zip":
		err = extractRar(src, dst)
	default:
		return fmt.Errorf("%s is not a supported archive: %s", src, srcType)
	}

	return err
}

func extractRar(src, dst string) error {
	path, err := exec.LookPath("rar")
	if err == nil {
		err := os.MkdirAll(dst, 0755)
		if err != nil {
			return err
		}

		cmd := exec.Command(path, "x", src, "*", dst)
		cmd.Dir = dst

		return cmd.Run()
	}
	return err
}
