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
//    "strings"
//    "io"
//    "os"
//    "lib/dt"
    "lib/cli"
)

func main() {

    // Define app commands based on CLI args
    commands := cli.Commands
    //var command cli.Command
    for {
        if command := commands.Shift(); "" != command.Name {
            fmt.Printf("\n\n----------------------\nname: %v; opts: %v, flags, %v\n----------------------------\n\n", command.Name, command.Opts, command.Flags)
        } else {
            break;
        }
    }









//fmt.Printf("\n\n----------------------\n%v\n----------------------------\n\n", commands.Shift())
//fmt.Printf("\n\n----------------------\n%v\n----------------------------\n\n", commands.Shift())
//fmt.Printf("\n\n----------------------\n%v\n----------------------------\n\n", commands.Shift())

//    for k, v := range commands {
//        fmt.Printf("command: '%v';\n", k)
//        fmt.Printf("    options: '%v';\n", v.Opts)
//        fmt.Printf("    flags: '%v';\n\n\n", v.Flags)
//    }
//
//    fmt.Printf("%v\n", commands["woot"].HasOpt("guido"))
//    fmt.Printf("%v\n", commands["woot"].GetOpt("guido"))
//    fmt.Printf("%v\n", commands["asdf"].HasOpt("guido"))
//    fmt.Printf("%v\n", commands["asdf"].GetOpt("guido"))






//    //a := make(map[int]string)
//
//    a := strings.Split("foo", "=")
//
//    fmt.Printf("k: %v\n", a[0])
//    fmt.Printf("v: %v\n", a[1])

//    for k, v := range a {
//    }

    //a[0] = "zero"
    //a[1] = "one"
    //fmt.Printf("0: %s\n", a[0])
    //fmt.Printf("1: %s\n", a[1])
    //fmt.Printf("len: %v\n", len(a))
    //fmt.Printf("cap: %v\n", cap(a))

//    fmt.Printf("Arguments:")
//    for _, arg := range os.Args {
//        fmt.Printf("    - %s", arg)
//    }

//    flag.Usage = func() {
//        fmt.Printf("Usage of %s:\n", os.Args[0])
//        fmt.Printf("    cat file1 file2 ...\n")
//        flag.PrintDefaults()
//    }
//
//    flag.Parse()
//    if flag.NArg() == 0 {
//        flag.Usage()
//        os.Exit(1)
//    }
//
//    for _, fn := range flag.Args() {
//        f, err := os.Open(fn);
//        if err != nil {
//            panic(err)
//        }
//        _, err = io.Copy(os.Stdout, f)
//        if err != nil {
//            panic(err)
//        }
//    }
}






