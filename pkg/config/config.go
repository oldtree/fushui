package config

import (
	"encoding/json"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
)

var fg = new(Config)

type Config struct {
	Address      string
	Port         string
	Database     string
	RedisAddress string
}

func (c *Config) LoadConfig(configfilepath string) (bool, error) {
	data, err := ioutil.ReadFile(configfilepath)
	if err != nil {
		log.Errorf("read config file error : ", err.Error())
		return false, err
	}
	err = json.Unmarshal(data, c)
	if err != nil {
		log.Errorf("json parse file error : ", err.Error())
		return false, err
	}
	log.Infof("load config file success")
	return true, nil
}
