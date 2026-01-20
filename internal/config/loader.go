package config

import (
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type EnvMap struct {
	Mappings map[string]string `mapstructure:"mappings,omitempty"`
}

type SecretMap struct {
	Dir      string            `mapstructure:"dir,omitempty"`
	Mappings map[string]string `mapstructure:"mappings,omitempty"`
}

func readConfig(logger *slog.Logger, defaultConfigValues *viper.Viper, name string, ext string, dirs ...string) (*viper.Viper, error) {
	logger.Info("Reading the configuration file", "file", fmt.Sprintf("%s.%s", name, ext), "dirs", fmt.Sprintf("%v", dirs))

	configValues := viper.New()

	if defaultConfigValues != nil {
		// set the default values
		for _, key := range defaultConfigValues.AllKeys() {
			configValues.SetDefault(key, defaultConfigValues.Get(key))
		}
	}

	configValues.SetConfigName(name) // name of config file (without extension)
	configValues.SetConfigType(ext)  // REQUIRED if the config file does not have the extension in the name
	for _, dir := range dirs {
		configValues.AddConfigPath(dir)
	}
	err := configValues.ReadInConfig() // Find and read the config file

	if err != nil {
		logger.Error("Failed to read the configuration file", "file", fmt.Sprintf("%s.%s", name, ext), "dirs", fmt.Sprintf("%v", dirs), "error", err.Error())
	} else {
		logger.Info("Read the configuration file", "file", configValues.ConfigFileUsed())
	}

	return configValues, err
}

func LoadConfig(logger *slog.Logger) (*Config, error) {
	// first load the server.yaml as the default config (the server.yaml from cmd/eval_hub)
	defaultConfigValues, err := readConfig(logger, nil, "server", "yaml", "config", "./cmd/eval_hub")
	if err != nil {
		return nil, err
	}

	// now load the cluster config if found
	configValues, err := readConfig(logger, defaultConfigValues, "config", "yaml", ".", "..")
	// for now we ignre this error because there is no extra config when running locally
	// if err != nil {
	//	 return nil, err
	// }

	// set up the secrets from the secrets directory
	secrets := SecretMap{}
	if err := configValues.Unmarshal(&secrets); err != nil {
		return nil, err
	}
	if secrets.Dir != "" {
		// check that the secrets directory exists
		if _, err := os.Stat(secrets.Dir); !os.IsNotExist(err) {
			for fieldName, value := range secrets.Mappings {
				secret := getSecret(secrets.Dir, value)
				if secret != "" {
					configValues.Set(fieldName, secret)
				}
			}
		}
	}
	// set up the environment variable mappings
	envMappings := EnvMap{}
	if err := configValues.Unmarshal(&envMappings); err != nil {
		return nil, err
	}
	for fieldName, values := range envMappings.Mappings {
		envNames := strings.Split(values, ",")
		elems := make([]string, 0, len(envNames)+1)
		elems = append(elems, fieldName)
		elems = append(elems, envNames...)
		configValues.BindEnv(elems...)
	}

	conf := Config{}
	if err := configValues.Unmarshal(&conf); err != nil {
		return nil, err
	}
	return &conf, nil
}

func getSecret(secretsDir string, secretName string) string {
	secret, err := os.ReadFile(fmt.Sprintf("%s/%s", secretsDir, secretName))
	if err != nil {
		return ""
	}
	return string(secret)
}
