package main

import (
    "fmt"
    "log"
    "os"
    "os/exec"
)

func main() {
    fmt.Println("Its kicking off...")

    if len(os.Args) < 2 {
        // If argument is not provided quit
        log.Fatalln("url not provided")
    }
    url := os.Args[1] // URL

    // cmd := exec.Command("curl", "-O", url)
    cmd := exec.Command("curl", "-o", "output.txt", url)
    cmd.Run()
}

