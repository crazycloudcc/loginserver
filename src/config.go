package main

import (
	"cherry/base"
	"cherry/base/config"
	"cherry/dbproxy"
	"cherry/nethttp"
	"cherry/nettcp"
	"cherry/netwebsocket"
)

// ConfigFile TODO
type ConfigFile struct {
	AppConf struct {
		ID   int32
		Name string
	}
	HTTPConf        nethttp.Config
	HTTPSConf       nethttp.Config
	WSConf          netwebsocket.Config
	TCPConf         nettcp.Config // tcp
	RedisConfRemote dbproxy.RedisConfig
	RedisConfLocal  dbproxy.RedisConfig
}

var configData *ConfigFile

// LoadConfigFile TODO
func loadConfigFile() *ConfigFile {
	if configData == nil {
		configData = new(ConfigFile)
		err := config.Read("./config/conf.json", "json", configData)
		if err != nil {
			base.LogFatal("LoadFile error [conf.json]: ", err)
		}
	}
	return configData
}
