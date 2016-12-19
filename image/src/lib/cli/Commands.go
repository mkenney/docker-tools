/*
Package cli provides a structure for multiple CLI arguments and their options,
flags and additional data.

On init, an instance of the CommandsList type is exported as Commands for
immediate use by the application. New() is provided if you ever need to
re-construct the commands stack. Shift() is provided to shift the next command
off of the stack and return and updated copy of the stack
*/
package cli

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
)

/*
Commands is an array of Command instances.
*/
var Commands CommandList

/*
CommandList is an ordered slice of Command structs

A Command is a named collection of command line options and flags groupped
by the preceeding CLI argument.

    - A command is defined as a CLI input that does not begin with a dash ('-').
    - An option is a CLI input that begins with '--' and is expected to be
      a 'key=value' formatted string. '--' results in the option '--'.
    - A flag is a CLI input that begins with '-' (including the argument '--')
	  and is is a boolean (was either passed in or not)
        An input that begins with '---' will be parsed into an option called
		'-input'

Command.Name  - command name
Command.Args  - array of all other argument strings
Command.Flags - map[flag]true
Command.Opts  - map[option]value

available Command methds:
    command.HasFlag(flag string) bool
    command.HasOpt(name string)  bool
    command.GetOpt(name string)  string

You can rebuild the global reference by running
	`cli.Commands = cli.New()`
*/
type CommandList []Command

/*
Shift returns a copy of the command stack with the first argument shifted out
*/
func (cmds CommandList) Shift() (retptr CommandList, retcmd Command, reterr error) {
	retptr = make(CommandList, len(cmds)-1)

	if 0 == len(cmds) {
		reterr = errors.New("No commands found")
	} else {
		retcmd = cmds[0]
		for a := 1; a < len(cmds); a++ {
			retptr[a-1] = cmds[a]
		}
	}

	return
}

/*
New constructs a pointer to a CommandList instance
*/
func New() (retval CommandList) {
	args := os.Args

	// args[0] includes the absolute path
	command := ""
	cmdargs := make([]string, 0, 100)
	flags := make(map[string]bool)
	opts := make(map[string]string)

	for _, arg := range args {
		fmt.Printf("arg: %v;\n", arg)

		// Is a command name
		// Command names consist of letters, numbers, underscores, hyphens,
		// periods and forward-slashes, but do not begin with a hyphen.
		re := regexp.MustCompile("^[a-zA-Z0-9_\\-\\./]+$")
		if !strings.HasPrefix(arg, "-") && re.Match([]byte(arg)) {
			// if it's a path, only store the basename
			if strings.Contains(arg, "/") {
				pieces := strings.Split(arg, "/")
				arg = pieces[len(pieces)-1]
			}
			// Update on the second pass in order to store the correct collection
			// of options, flags and arguments with the command name
			if "" != command {
				cmd := Command{command, cmdargs[:], flags, opts}
				retval = append(retval, cmd)
				cmdargs = make([]string, 0, 100)
				opts = make(map[string]string)
				flags = make(map[string]bool)
			}
			command = arg

			// begins with --, is an opt
			// is only '--', is the opt '--'
		} else if strings.HasPrefix(arg, "--") {
			opt := strings.Split(arg[2:], "=")
			if "" == opt[0] {
				opt[0] = "--"
			}
			if 1 == len(opt) {
				opts[opt[0]] = ""
			} else {
				opts[opt[0]] = opt[1]
			}

			// begins with -, is a flag
			// is only '-', is the flag '-'
		} else if strings.HasPrefix(arg, "-") {
			if "" == arg[1:] {
				arg = "--"
			}
			flags[arg[1:]] = true

			// everything else is an "argument"
		} else {
			cmdargs = append(cmdargs, arg)
		}
	}
	cmd := Command{command, cmdargs[:], flags, opts}
	retval = append(retval, cmd)

	return
}

func init() {
	Commands = New()
}
