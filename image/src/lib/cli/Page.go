/*
Package cli provides a structure for multiple CLI arguments and their options,
flags and additional data.

Types that implement the Renderer interface should return text output formatted
for your targeted terminal interface (xterm, rxvt, dumb, etc.).
*/
package cli

import (
    "fmt"
    "os"
    "os/exec"

	"github.com/golang/glog"
)

func Page(str string) {
	pipestdin, pipestdout, err := os.Pipe()
	if err != nil {
		panic("Could not create pipe")
	}

	stdout := os.Stdout
	os.Stdout = pipestdout

	pager := exec.Command("less", "-r")
	pager.Stdin = pipestdin
	pager.Stdout = stdout // the pager uses the original stdout, not the pipe...
	pager.Stderr = os.Stderr

	defer func() {
		pipestdout.Close()
		err := pager.Run()
		os.Stdout = stdout
		if err != nil {
			glog.Fatalf("%v", os.Stderr)
			glog.Fatalf("%s", err)
		}
	}()

	fmt.Println("\n\n" + str)

}
