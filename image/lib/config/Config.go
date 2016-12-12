
package config

import (
    "os"
)

const DEFAULT_CONFIG_DIR = "/usr/local/docker-tools"
const DEFAULT_PREFIX_DIR = "/usr/local/bin"

/**
 * Represents all tool recipes available to the system
 */
type Config struct {
    Path string
    Prefix string
}

/**
 *
 */
func New() *Config {
    config := new(Config)

    config.Path   = map[bool]string{true: os.Getenv("DOCKER_TOOLS_CONFIG_DIR"), false: DEFAULT_CONFIG_DIR} ["" != os.Getenv("DOCKER_TOOLS_CONFIG_DIR")]
    config.Prefix = map[bool]string{true: os.Getenv("DOCKER_TOOLS_PREFIX_DIR"), false: DEFAULT_PREFIX_DIR} ["" != os.Getenv("DOCKER_TOOLS_PREFIX_DIR")]

    return config
}
