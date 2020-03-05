package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"sync"
)

type Configs map[string]json.RawMessage
var configPath string = "config.json"


type MainConfig struct {
	Port string `json:"port"`
}


var conf *MainConfig
var confs Configs

var instanceOnce sync.Once

//从配置文件中载入json字符串
func LoadConfig(path string) (Configs, *MainConfig) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		log.Panicln("load config conf failed: ", err)
	}
	mainConfig := &MainConfig{}
	err = json.Unmarshal(buf, mainConfig)
	if err != nil {
		log.Panicln("decode config file failed:", string(buf), err)
	}
	allConfigs := make(Configs, 0)
	err = json.Unmarshal(buf, &allConfigs)
	if err != nil {
		log.Panicln("decode config file failed:", string(buf), err)
	}

	return allConfigs, mainConfig
}

//初始化 可以运行多次
func SetConfig(path string) {
	allConfigs, mainConfig := LoadConfig(path)
	configPath = path
	conf = mainConfig
	confs = allConfigs
}

// 初始化，只能运行一次
func Init(path string) *MainConfig {
	if conf != nil && path != configPath {
		log.Printf("the config is already initialized, oldPath=%s, path=%s", configPath, path)
	}
	instanceOnce.Do(func() {
		allConfigs, mainConfig := LoadConfig(path)
		configPath = path
		conf = mainConfig
		confs = allConfigs
	})

	return conf
}

//初始化配置文件 为 struct 格式
func Instance() *MainConfig {
	if conf == nil {
		Init(configPath)
	}
	return conf
}


//初始化配置文件 为 map格式
func AllConfig() Configs {
	if conf == nil {
		Init(configPath)
	}
	return confs
}

//获取配置文件路径
func ConfigPath() string {
	return configPath
}

//根据key获取对应的值，如果值为struct，则继续反序列化
func (cfg Configs) GetConfig(key string, config interface{}) error {
	c, ok := cfg[key]
	if ok {
		return json.Unmarshal(c, config)
	} else {
		return fmt.Errorf("fail to get cfg with key: %s", key)
	}
}
