Run with:
```go
go run main.go camera
```
for realtime (there is a small delay) detection from camera  
or
```go
go run main.go path-to-image
```
for detection from an image

If your camera is not found, try changing the device parameter in gocv.VideoCaptureDevice (line 37) in detectFromCamera func from 0 to 1  
The gocv library should download automatically on startup, if not, use this command:
```go
go get -u -d gocv.io/x/gocv
```