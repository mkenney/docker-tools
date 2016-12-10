package cli

import (
//    "fmt"
    "os"
    "strings"
)

//////////////////////////////////////////////////////////////////////////////
// Command
//////////////////////////////////////////////////////////////////////////////
/**
 * Command properties
 * A command can have options and flags. Flags, if available, should be true,
 * otherwise a flag check should return nil
 *
 * available methds:
 *     command.HasFlag(flag string) bool:
 *     command.HasOpt(name string) bool:
 *     command.GetOpt(name string) string:
 */
type Command struct {
    Name string
    Opts map[string]string
    Flags map[string]bool
}

    /**
     * Test to see if a flag has been passed
     */
    func (this Command) HasFlag(flag string) bool {
        var ret_val bool
        if _, ok := this.Flags[flag]; ok {
            ret_val = true
        }
        return ret_val
    }
    /**
     * Test to see if an option has been passed
     */
    func (this Command) HasOpt(opt string) bool {
        var ret_val bool = false
        if _, ok := this.Opts[opt]; ok {
            ret_val = true
        }
        return ret_val
    }
    /**
     * Get an option value
     */
    func (this Command) GetOpt(opt string) string {
        var ret_val string
        if val, ok := this.Opts[opt]; ok {
            ret_val = val
        }
        return ret_val
    }

//////////////////////////////////////////////////////////////////////////////
// cliCommands
//////////////////////////////////////////////////////////////////////////////
/**
 * cliCommands is an ordered collection of Command structs
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
type cliCommands map[string]Command

    /**
     * Shift the first argument out of the stack
     */
    func (this cliCommands) Shift() Command {
        var ret_val Command
        for _, val := range this {
            ret_val = val
            delete(this, val.Name)
            break
        }
        return ret_val
    }

    /**
     * Commands instance factory...
     */
    func (self cliCommands) Factory(args []string) cliCommands {
        ret_val := make(cliCommands)

        // args[0] includes the absolute path
        arg_parts := strings.Split(args[0], "/")
        var command string = arg_parts[len(arg_parts) - 1]

        flags := make(map[string]bool)
        opts  := make(map[string]string)

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
                ret_val[command] = Command{command, opts, flags}
                opts = make(map[string]string)
                flags = make(map[string]bool)

                command = arg
            }
        }

        ret_val[command] = Command{command, opts, flags}


        return ret_val;
    }


var args [] string = os.Args
var Commands cliCommands = cliCommands.Factory(args)

