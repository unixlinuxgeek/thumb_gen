package thumb_gen

import (
	"errors"
	"os"
	"strconv"
	"testing"
	"time"
)

func TestGen(t *testing.T) {
	tn := time.Now()
	u := tn.Unix()
	tp := strconv.FormatInt(u, 10)
	err := Gen("./test_data/in.mp4", "./test_data/out_"+tp+".png", "00:00:02")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat("./test_data/out_" + tp + ".png"); errors.Is(err, os.ErrNotExist) {
		t.Fatal(err)
	}
}
