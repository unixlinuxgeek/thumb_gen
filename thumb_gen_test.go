package thumb_gen

import (
	"strconv"
	"testing"
	"time"
)

func TestGen(t *testing.T) {
	tn := time.Now()
	u := tn.Unix()
	err := Gen("./test_data/in.mp4", "./test_data/out_"+strconv.FormatInt(u, 10)+".png", "00:01:00")
	if err != nil {
		t.Fatal(err)
	}
}
