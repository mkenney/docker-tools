
package config

import (
    "os"
    "strings"
)

const DEFAULT_CONFIG_DIR = "/usr/local/docker-tools"
const DEFAULT_PREFIX_DIR = "/usr/local/bin"

/**
 * Represents all tool recipes available to the system
 */
type ConfigStruct struct {
    HostPath []string
    HostConfPath string
    HostHome string
    ConfPath string
    DefaultPrefix string
    ToolsVersion string
}

/**
 *
 */
func NewConfig() *ConfigStruct {
    host_path := map[bool]string{true: os.Getenv("HOST_PATH"), false: "/host"+DEFAULT_PREFIX_DIR} ["" != os.Getenv("HOST_PATH")]

    config := new(ConfigStruct)
    config.ConfPath      = "/usr/local/docker-tools"

    config.HostPath      = strings.Split(host_path, ":")
    config.HostHome      = map[bool]string{true: os.Getenv("HOME"), false: ""} ["" != os.Getenv("HOME")]
    config.HostConfPath  = map[bool]string{true: os.Getenv("DOCKER_TOOLS_CONFIG_DIR"), false: DEFAULT_CONFIG_DIR} ["" != os.Getenv("DOCKER_TOOLS_CONFIG_DIR")]
    config.DefaultPrefix = map[bool]string{true: os.Getenv("DOCKER_TOOLS_PREFIX_DIR"), false: DEFAULT_PREFIX_DIR} ["" != os.Getenv("DOCKER_TOOLS_PREFIX_DIR")]
    config.ToolsVersion  = "v0.0.1"

    return config
}


var Config *ConfigStruct
func init() {
    Config = NewConfig()
}
