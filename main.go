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
    "flag"
    "fmt"
    "io"
    "os"
)

func main() {
    flag.Usage = func() {
        fmt.Printf("Usage of %s:\n", os.Args[0])
        fmt.Printf("    cat file1 file2 ...\n")
        flag.PrintDefaults()
    }

    flag.Parse()
    if flag.NArg() == 0 {
        flag.Usage()
        os.Exit(1)
    }

    for _, fn := range flag.Args() {
        f, err := os.Open(fn);
        if err != nil {
            panic(err)
        }
        _, err = io.Copy(os.Stdout, f)
        if err != nil {
            panic(err)
        }
    }
}






