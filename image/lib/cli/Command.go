package cli

import ()

type CommandInterface interface {
    GetOpt(opt string) string
    HasOpt(opt string) bool
    HasFlag(flag string) bool
}

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

    /**
     * The name of the command
     */
    Name string

    /**
     * A key/value map of command-line options. An "option" is defined as any
     * argument that begins with "--". An option without a value or an empty
     * value defaults to an empty string.
     */
    Opts map[string]string

    /**
     * A key/value map of
     */
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
