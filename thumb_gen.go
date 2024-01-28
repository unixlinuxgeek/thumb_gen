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
	arg2 := "-ss"
	arg3 := time
	arg4 := "-i"
	arg5 := vidName
	arg6 := "-frames:v"
	arg7 := "1"
	arg8 := outImg

	cmd := exec.Command(app, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
