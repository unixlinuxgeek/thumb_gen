package thumb_gen

import (
	"errors"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	path, err := exec.LookPath("ffmpeg")
	if errors.Is(err, exec.ErrDot) {
		err = nil
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(path)
	//use(path)
}
