package thumb_gen

import (
	"errors"
	"fmt"
	"os/exec"
)

func Installed() bool {
	path, err := exec.LookPath("ffmpeg")
	if errors.Is(err, exec.ErrDot) {
		err = nil
	}
	if err != nil {
		return false
	}
	fmt.Println(path)
	return true
}
