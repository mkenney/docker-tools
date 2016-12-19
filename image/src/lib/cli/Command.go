package cli

/*
CommandInterface defines the interface for accessing the data in the the Command
struct
*/
type CommandInterface interface {
	GetOpt(opt string) string
	HasOpt(opt string) bool
	HasFlag(flag string) bool
}

/*
Command defines a single, complete command

A command is defined as a single statement (a word that begins with an
alphanumeric character and ends with a space, EOL or other whitespace) followed
by options (words beginning with '--', optionally with a value assignment) or
flags (words beginning with '-'). The special case '--' followed by whitespace
is translated as the flag '--' for use as a separator (ala git, etc.).
*/
type Command struct {

	/*
	   the command that was passed
	*/
	Name string

	/*
	   A list of all remaining arguments
	*/
	Args []string

	/*
		A key/value map of each flag that was passed and the value `true`
	*/
	Flags map[string]bool

	/*
	   A key/value map of command-line options. An "option" is defined as any
	   argument that begins with "--". An option without a value or an empty
	   value defaults to an empty string.
	*/
	Opts map[string]string
}

/*
GetOpt returns an option value
If no value was set or the option ended in '=', an emtpy string is returned
*/
func (cmd Command) GetOpt(opt string) string {
	var retval string
	if val, ok := cmd.Opts[opt]; ok {
		retval = val
	}
	return retval
}

/*
HasArg tests to see if a flag has been passed to this command
*/
func (cmd Command) HasArg(arg string) (retval bool) {
	for _, val := range cmd.Args {
		if val == arg {
			retval = true
			break
		}
	}
	return
}

/*
HasFlag tests to see if a flag has been passed to this command
*/
func (cmd Command) HasFlag(flag string) bool {
	var retval bool
	if _, ok := cmd.Flags[flag]; ok {
		retval = true
	}
	return retval
}

/*
HasOpt tests to see if an option has been passed to this command
*/
func (cmd Command) HasOpt(opt string) bool {
	retval := false
	if _, ok := cmd.Opts[opt]; ok {
		retval = true
	}
	return retval
}
