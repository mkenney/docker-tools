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
DockerToolsVersion is the curremt semantic version
*/
const DockerToolsVersion = "0.0.0"

/*
DockerToolsRegistry is the location of the tool registry file
*/
const DockerToolsRegistry = "/go/src/docker-tools/registry.yml"

/*
DockerToolsConfigDir is the configuration directory
*/
const DockerToolsConfigDir = "/usr/local/docker-tools"

/*
DockerToolsTemplateDir is the template file location
*/
const DockerToolsTemplateDir = "/go/src/docker-tools/src/lib/templates"

/*
DockerToolsDefaultPrefixDir is the default installation directory for scripts on the host
system
*/
const DockerToolsDefaultPrefixDir = "/usr/local/bin"

/*
ValueData is a struct that represents all tool recipes available to the system
*/
var (
	HostPath      []string
	HostConfPath  string
	HostHome      string
	ConfPath      string
	DefaultPrefix string
)

/*
New initializes and returns a pointer to a ValueData struct containing current
configuration values for the docker-tools runtime
*/
func init() {
	hostPath := map[bool]string{true: os.Getenv("HOST_PATH"), false: "/host" + DockerToolsDefaultPrefixDir}["" != os.Getenv("HOST_PATH")]

	ConfPath = "/usr/local/docker-tools"
	HostPath = strings.Split(hostPath, ":")
	HostHome = map[bool]string{true: os.Getenv("HOME"), false: ""}["" != os.Getenv("HOME")]
	HostConfPath = map[bool]string{true: os.Getenv("DOCKER_TOOLS_CONFIG_DIR"), false: DockerToolsConfigDir}["" != os.Getenv("DOCKER_TOOLS_CONFIG_DIR")]
	DefaultPrefix = map[bool]string{true: os.Getenv("DOCKER_TOOLS_PREFIX_DIR"), false: DockerToolsDefaultPrefixDir}["" != os.Getenv("DOCKER_TOOLS_PREFIX_DIR")]
}
