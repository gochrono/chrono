package chronolib

import (
    "os"
    "github.com/kirsle/configdir"
    "path/filepath"
)

// ChronoAppConf is the name of the app's data directory
const ChronoAppConf = "chrono"

// ChronoConfDirEnvName is name of the environment variable used to manually set the config directory
const ChronoConfDirEnvName = "CHRONO_CONFIG_DIR"

// GetAppFilePath returns a file's path in the config directory through either an environment variable or the default path
func GetAppFilePath(fileName string, customConfDir string) string {
    var appConfDir = configdir.LocalConfig(ChronoAppConf)
    if os.Getenv(ChronoConfDirEnvName) != "" {
        appConfDir = os.Getenv(ChronoConfDirEnvName)
    }
    if customConfDir != "" {
        appConfDir = customConfDir
    }
    err := configdir.MakePath(appConfDir)
    if err != nil { panic(err) }
    return filepath.Join(appConfDir, fileName)
}

