/*
Package cli provides a structure for multiple CLI arguments and their options,
flags and additional data.

Types that implement the Renderer interface should return text output formatted
for your targeted terminal interface (xterm, rxvt, dumb, etc.).
*/
package cli

/*
Renderer implements a Render method for a Type to generate CLI output
*/
type Renderer interface {
	Render(data, options map[string]string) string
}
