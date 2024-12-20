package main

import (
    "fmt"
    "os"
    "os/exec"
    "time"
)

func main() {
    videoURL := "https://www.youtube.com/shorts/LW1-ZfkqEJc"
    outputFile := "output.mp4"

    start := time.Now()

    cmd := exec.Command("yt-dlp", "-o", outputFile, videoURL)

    
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr


    err := cmd.Run()
    if err != nil {
        fmt.Println("Error downloading video:", err)
        return
    }

    elapsed := time.Since(start)
    fmt.Println("Video downloaded successfully as", outputFile)
    fmt.Printf("Execution time: %s\n", elapsed)
}

