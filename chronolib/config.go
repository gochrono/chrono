package chronolib

import (
    "os"
    "github.com/kirsle/configdir"
    "path/filepath"
)

const ChronoAppConf = "chrono"
const ChronoConfDirEnvName = "CHRONO_CONFIG_DIR"

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

