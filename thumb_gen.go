// Capture 1 frame from video

package thumb_gen

import (
	"log"
	"os"
	"os/exec"
)

func Gen(vidName string, outImg, time string) error {
	f, err := os.OpenFile(vidName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 777)
	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}

	app := "./bin/ffmpeg-amd64-static/ffmpeg"
	arg1 := "-y"
	arg2 := "-i"
	arg3 := vidName
	arg4 := "-frames:v"
	arg5 := "1"
	arg6 := outImg

	cmd := exec.Command(app, arg1, arg2, arg3, arg4, arg5, arg6)
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
