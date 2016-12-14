//package main
//
//import (
//    "flag"
//    "fmt"
//    "io"
//    "os"
//)
//
//func main() {
//    var port int
//
//    // Program Name is always the first (implicit) argument
//    cmd := os.Args[0]
//    fmt.Printf("Program Name: %s\n", cmd)
//
//
//    argCount := len(os.Args[1:])
//    fmt.Printf("Total Arguments (excluding program name): %d\n", argCount)
//
//
//    for i, a := range os.Args[1:] {
//        fmt.Printf("Argument %d is %s\n", i+1, a)
//    }
//
//
//    flag.IntVar(&port, "p", 8000, "specify port to use.  defaults to 8000.")
//    flag.Parse()
//    fmt.Printf("port = %d", port)
//    fmt.Printf("other args: %+v\n", flag.Args())
//
//
//}

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
