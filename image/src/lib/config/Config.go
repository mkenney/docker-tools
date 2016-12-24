/*
Package config provides a struct and methds for interacting with the
docker-tools configuration file
*/
package config

import (
	"os"
	"strings"
)

/*
DefaultConfigDir is the default configuration directory inside the container
*/
const DefaultConfigDir = "/usr/local/docker-tools"

/*
DefaultPrefixDir is the default installation directory for scripts on the host
system
*/
const DefaultPrefixDir = "/usr/local/bin"

/*
ValueData is a struct that represents all tool recipes available to the system
*/
type ValueData struct {
	HostPath      []string
	HostConfPath  string
	HostHome      string
	ConfPath      string
	DefaultPrefix string
	ToolsVersion  string
}

/*
New initializes and returns a pointer to a ValueData struct containing current
configuration values for the docker-tools runtime
*/
func New() *ValueData {
	hostPath := map[bool]string{true: os.Getenv("HOST_PATH"), false: "/host" + DefaultPrefixDir}["" != os.Getenv("HOSData_PATH")]

	config := new(ValueData)
	config.ConfPath = "/usr/local/docker-tools"

	config.HostPath = strings.Split(hostPath, ":")
	config.HostHome = map[bool]string{true: os.Getenv("HOME"), false: ""}["" != os.Getenv("HOME")]
	config.HostConfPath = map[bool]string{true: os.Getenv("DOCKER_TOOLS_CONFIG_DIR"), false: DefaultConfigDir}["" != os.Getenv("DOCKER_TOOLS_CONFIG_DIR")]
	config.DefaultPrefix = map[bool]string{true: os.Getenv("DOCKER_TOOLS_PREFIX_DIR"), false: DefaultPrefixDir}["" != os.Getenv("DOCKER_TOOLS_PREFIX_DIR")]
	config.ToolsVersion = "v0.0.1"

	return config
}

/*
Values exports the docker-tools configuration data
*/
var Values *ValueData

func init() {
	Values = New()
}
