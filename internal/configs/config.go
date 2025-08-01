package configs

import (
	"github.com/spf13/viper"
)

var config *Config

type option struct {
	configFolders  []string
	configFilename string
	configType     string
}

func Init(opts ...Option) error {
	opt := &option{
		configFolders:  getDefaultConfigFolder(),
		configFilename: getDefaultConfigFilename(),
		configType:     getDefaultConfigType(),
	}

	// override option (jika ingin mengubah config file, jika tidak, maka akan menggunakan default)
	for _, optFunc := range opts {
		// inject custom path ke opt
		optFunc(opt)
	}

	for _, configFolder := range opt.configFolders {
		// add config path to viper
		viper.AddConfigPath(configFolder)
	}

	viper.SetConfigName(opt.configFilename)
	viper.SetConfigType(opt.configType)
	viper.AutomaticEnv()

	config = new(Config)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return viper.Unmarshal(&config)
}

type Option func(*option)

func getDefaultConfigFolder() []string {
	return []string{"./configs"}
}

func getDefaultConfigFilename() string {
	return "config"
}

func getDefaultConfigType() string {
	return "yaml"
}

func WithConfigFolder(configFolders []string) Option {
	return func(opt *option) {
		opt.configFolders = configFolders
	}
}

func WithConfigFilename(configFilename string) Option {
	return func(opt *option) {
		opt.configFilename = configFilename
	}
}

func WithConfigType(configType string) Option {
	return func(opt *option) {
		opt.configType = configType
	}
}

func Get() *Config {
	if config == nil {
		config = &Config{}
	}
	return config
}
