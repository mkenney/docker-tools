/**
 * cli.Commands is a global pointer to the current commands index
 */

package cli

import (
    "errors"
//    "fmt"
    "os"
    "strings"
)

/**
 * CliCommands is an unordered collection of Command structs. An ordered index
 * (CliCommandsIndex) is provided.
 *
 * A Command is a named collection of command line options and flags groupped
 * by the preceeding CLI argument.
 *
 *     - An argument is defined as a CLI input that does not begin with a dash ('-').
 *     - An option is a CLI input that begins with '--' (excluding '---') and is expected to be
 *       'key=value' strings
 *     - A flag is a CLI input that begins with '-' (including '---') and is is a boolean (was
 *       either passed in or not)
 *         An input that begins with '---' will be parsed into a flag '--%restofinput'
 *
 * Command.Name  - command name
 * Command.Opts  - map of option:value
 * Command.flags - map of flag:true
 *
 * available Command methds:
 *     command.HasFlag(flag string) bool:
 *     command.HasOpt(name string) bool:
 *     command.GetOpt(name string) string:
 */
type CliCommands [100]Command

/**
 * Return a copy of the command stack with the first argument shifted out and
 * a new pointer to an updated command stack.
 */
func (this *CliCommands) Shift() (*CliCommands, Command, error) {
    ret_ptr := new(CliCommands)
    var ret_cmd Command
    var ret_err error

    if 0 == len(this) {
        ret_err = errors.New("No commands found")

    } else {
        ret_cmd = this[0]
        for a := 1; a < len(this); a++ {
            ret_ptr[a - 1] = this[a]
        }
    }

    return ret_ptr, ret_cmd, ret_err
}

/**
 * Commands instance factory...
 */
func cliCommandsFactory(args []string) (*CliCommands) {
    ret_val := new(CliCommands)
    var index int

    // args[0] includes the absolute path
    arg_parts := strings.Split(args[0], "/")
    var command string = arg_parts[len(arg_parts) - 1]

    flags := make(map[string]bool)
    opts  := make(map[string]string)

    index = 0
    for _, arg := range args[1:] {

        // begins with -- is an opt
        if strings.HasPrefix(arg, "--") && !strings.HasPrefix(arg, "---") {
            opt := strings.Split(arg[2:], "=")
            if 1 == len(opt) {
                opts[opt[0]] = ""
            } else {
                opts[opt[0]] = opt[1]
            }

        // begins with - or --- is a flag
        } else if strings.HasPrefix(arg, "---"); strings.HasPrefix(arg, "-") {
            flags[arg[1:]] = true

        // is a command
        } else {
            ret_val[index] = Command{command, opts, flags}
            index++

            opts = make(map[string]string)
            flags = make(map[string]bool)
            command = arg
        }
    }

    ret_val[index] = Command{command, opts, flags}

    return ret_val
}

var (
    Commands *CliCommands
)

func init() {
    Commands = cliCommandsFactory(os.Args)
}
