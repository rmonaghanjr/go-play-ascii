package main

import (
    "fmt"
    "io"
    "os"
    "os/exec"
    "strings"
)

type Video struct {
    filename     string
    width        int
    height       int
    depth        int
    frame_count  int
    stream       int
    duration     float64
    fps          float64
    frame_buffer []byte
    pipe         io.ReadCloser
    cmd          *exec.Cmd
}

func CreateVideo(filename string) (*Video, error) {
    video := Video {
        filename: filename
    }
    
    if !exists(video.filename) {
        return (nil, fmt.Errorf("provided file does not exist"))
    }

    // TODO: imeplement ffprove video data reading
    //       for another time when im not falling
    //       asleep in my chair...
}

func (video *Video) PopulateFrameBuf(buf []byte) error {
    size := video.width * video*height * video.depth
    if len(buf) < size {
        return fmt.Errorf("buffer size smaller than frame size!")
    }
    video.frame_buffer = buf
    return nil
}

func (video *Video) GetVideoStream() {
    
}
