package chronolib

import (
	"github.com/kirsle/configdir"
	jww "github.com/spf13/jwalterweatherman"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

// ChronoConfigFilename is the config's expected filename
const ChronoConfigFilename = "config.toml"

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
	if err != nil {
		panic(err)
	}
	return filepath.Join(appConfDir, fileName)
}

// GetCorrectConfigDirectory returns the correct config path where customConfigPath > Env Variable > Default value
func GetCorrectConfigDirectory(customConfigPath string) string {
	var appConfDir = configdir.LocalConfig(ChronoAppConf)

	if os.Getenv(ChronoConfDirEnvName) != "" {
		appConfDir = os.Getenv(ChronoConfDirEnvName)
	}
	if customConfigPath != "" {
		appConfDir = customConfigPath
	}
	err := configdir.MakePath(appConfDir)
	if err != nil {
		panic(err)
	}
	return appConfDir
}

// EnsureConfigDirExists makes sure the configDirectory path exists
func EnsureConfigDirExists(configDirectory string) error {
	return configdir.MakePath(configDirectory)
}

// GetConfig loads the a confi file if it exists and returns a ChronoConfig struct
func GetConfig(configDirectory string) ChronoConfig {
	storageType := viper.GetString("storage")
	if storageType != jsonStorageType && storageType != msgpackStorageType {
		jww.WARN.Printf("unknown storage type %s, using msgpack", storageType)
		storageType = "msgpack"
	}
	jww.INFO.Printf("using storage type %s", storageType)
	generalConfig := chronoGeneralConfig{}
	return ChronoConfig{configDirectory, storageType, generalConfig}
}

type chronoGeneralConfig struct {
	Storage string
}

// ChronoConfig contains the currently used config dir path as well as any configuration options
// stored in config.toml
type ChronoConfig struct {
	ConfigDir     string
	StorageType   string
	GeneralConfig chronoGeneralConfig
}
