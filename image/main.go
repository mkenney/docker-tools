package main

import (
    "fmt"
    "lib/dt"
)

func main() {
    fmt.Printf("func main()\n\n")

    // Initialize and execute DockerTools
    dt.New().Run()
}
