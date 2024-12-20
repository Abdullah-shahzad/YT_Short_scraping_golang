package main

import (
    "fmt"
    "os"
    "os/exec"
)

func main() {
    videoURL := "https://www.youtube.com/shorts/FsV4QtjSi1I"
    outputFile := "output.mp4"

    cmd := exec.Command("yt-dlp", "-o", outputFile, videoURL)

    
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr


    err := cmd.Run()
    if err != nil {
        fmt.Println("Error downloading video:", err)
        return
    }

    fmt.Println("Video downloaded successfully as", outputFile)
}

