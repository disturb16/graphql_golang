package settings

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"gopkg.in/yaml.v2"
)

// Config main configuration params
// Change this struct if you want to add or remove properties
// from conf.yaml
type Config struct {
	Port   int    `yaml:"Port"`
	APIKey string `yaml:"ApiKey"`
	UseGraphiQL bool `yaml:"UseGraphiQL"`
	DB     struct {
		Host     string `yaml:"Host"`
		Name     string `yaml:"Name"`
		Port     int    `yaml:"Port"`
		User     string `yaml:"User"`
		Password string `yaml:"Password"`
	} `yaml:"Database"`
}

// GetEnv returns env value, if empty, returns fallback value
func GetEnv(key string, fallback string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	return value
}

//Configuration returns main configuration object
func Configuration(configPath string) (*Config, error) {

	configYaml := &Config{}
	var err error
	var confFile []byte

	// check if conf-file is .yaml or .json
	confFile, err = ioutil.ReadFile(configPath + "conf.yaml")
	if err != nil {
		confFile, err = ioutil.ReadFile(configPath + "conf.json")
		if err != nil {
			log.Printf("Could not open conf.yaml or conf.json with path %s", configPath)
		}
	}

	//if file exists use its variables
	if err == nil {
		err = yaml.Unmarshal(confFile, &configYaml)
		if err != nil {
			return nil, err
		}
	}

	// Parse port numbers
	portStr := GetEnv("PORT", strconv.Itoa(configYaml.Port))
	dbPortStr := GetEnv("DB_PORT", strconv.Itoa(configYaml.DB.Port))

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, err
	}

	dbPort, err := strconv.Atoi(dbPortStr)
	if err != nil {
		return nil, err
	}

	//Search for OS environment and use conf-file as fallback
	config := &Config{}
	config.Port = port
	config.APIKey = GetEnv("APIKEY", configYaml.APIKey)
	config.DB.Name = GetEnv("DB_NAME", configYaml.DB.Name)
	config.DB.Host = GetEnv("DB_HOST", configYaml.DB.Host)
	config.DB.User = GetEnv("DB_USER", configYaml.DB.User)
	config.DB.Password = GetEnv("DB_PASS", configYaml.DB.Password)
	config.DB.Port = dbPort

	config.UseGraphiQL, err = strconv.ParseBool(GetEnv("UseGraphiQL", strconv.FormatBool(configYaml.UseGraphiQL)))
	if err != nil {
		return nil, err
	}

	return config, nil
}
