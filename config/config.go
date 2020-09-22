package config

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"io/ioutil"
	"os"
	"tickethub_service/util"
	"tickethub_service/util/log"
)

type Config struct {
	ServerHost string `json:"server_host"`
	ServerPort int    `json:"server_port"`

	DatabaseHost     string `json:"database_host"`
	DatabasePort     int    `json:"database_port"`
	DatabaseUser     string `json:"database_user"`
	DatabasePassword string `json:"database_password"`
	DatabaseName     string `json:"database_name"`
}

var (
	SystemConfig Config
	DB           *gorm.DB
)

func InitDB() error {
	mysqlUrl := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", SystemConfig.DatabaseUser,
		SystemConfig.DatabasePassword, SystemConfig.DatabaseHost, SystemConfig.DatabasePort, SystemConfig.DatabaseName)
	db, err := gorm.Open("mysql", mysqlUrl)
	if err != nil {
		log.Errorf("Connect database failed: %v", err)
		return err
	}

	DB = db
	return nil
}

func SetDefaultConfig() {
	SystemConfig.ServerHost = "0.0.0.0"
	SystemConfig.ServerPort = 8888

	SystemConfig.DatabaseHost = "127.1.0.0"
	SystemConfig.DatabasePort = 3306
	SystemConfig.DatabaseUser = "root"
	SystemConfig.DatabasePassword = "root"
}

func ReadConfigFromFile() error {
	configFile, err := os.Open(util.ConfigPath)
	defer configFile.Close()
	if err != nil {
		log.Errorf("Failed to open configFile[%s]", util.ConfigPath)
		SetDefaultConfig()
		return err
	}
	content, err := ioutil.ReadAll(configFile)
	if err != nil {
		log.Errorf("Failed to read config file content[%v]", configFile)
		SetDefaultConfig()
		return err
	}

	err = json.Unmarshal(content, &SystemConfig)
	if err != nil {
		log.Errorf("Failed to parse config file content[%v]", content)
		SetDefaultConfig()
		return err
	}

	log.Infof("config info :%v", SystemConfig)
	return nil
}
