### thumb_gen utility

Generate Thumb image from video.

Running tests:
```
$ go test
PASS
ok      github.com/unixlinuxgeek/thumb_gen      0.026s 
```


```
ffmpeg -ss 00:00:05 -i input -frames:v 1 -q:v 2 output.jpg
```
