package config

import (
	"os"
	"os/user"
	"path/filepath"

	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

type MeisterConfig struct {
	Version  string
	CacheDir string
	Meister  map[string]interface{}
}

const (
	DEFAULT_CONFIG_NAME = "meister"
)

var (
	DEFAULT_CONFIG_PATH string

	_logger *zerolog.Logger
)

func init() {
	usr, err := user.Current()
	if err != nil {
		panic("Unable to retrieve current user")
	}

	DEFAULT_CONFIG_PATH = filepath.Join(usr.HomeDir, ".meister")
}

func Init(logger *zerolog.Logger) (c *MeisterConfig) {
	_logger = logger

	viper.SetConfigName("meister")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(DEFAULT_CONFIG_PATH)
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			_logger.Warn().Msgf("%v", err)
			_logger.Info().Msg("Creating default configuration...")

			initDefaultConfig()
		} else {
			_logger.Fatal().Err(err).Msgf("Failed to load configuration")
		}
	}

	if err := viper.Unmarshal(&c); err != nil {
		_logger.Fatal().Err(err).Msg("Failed to initialize configuration")
	}

	return
}

func initDefaultConfig() {
	viper.SetDefault("version", "1.0")
	viper.SetDefault("meister.cache_dir", filepath.Join(DEFAULT_CONFIG_PATH, "cache"))
	viper.SetDefault("meister.stencil_packs.prompt_update", true)
	viper.SetDefault("meister.stencil_packs.auto_update", false)

	writeDefaultConfig()
}

func writeDefaultConfig() {
	// Create the default config path if it doesn't exist
	if _, err := os.Stat(DEFAULT_CONFIG_PATH); os.IsNotExist(err) {
		_logger.Info().Msgf("Config directory `%v` not found, creating it", DEFAULT_CONFIG_PATH)

		if err = os.Mkdir(DEFAULT_CONFIG_PATH, os.ModeDir|0700); err != nil {
			_logger.Warn().Err(err).Msg("Failed to create config directory")
		}
	}

	if err := viper.SafeWriteConfig(); err != nil {
		_logger.Warn().Err(err).Msg("Failed to write configuration file")
	}
}
