package assist

import (
	"cherry/base"
	"cherry/base/config"
	"cherry/dbproxy"
	"cherry/nethttp"
	"cherry/nettcp"
	"cherry/netwebsocket"
)

// ConfigFile 配置文件结构. (.json文件变更时, 需要进行对应修改.)
type ConfigFile struct {
	AppConf struct {
		ID    int32
		Group string
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
func LoadConfigFile(fn string) *ConfigFile {
	if configData == nil {
		configData = new(ConfigFile)
		err := config.Read(fn, "json", configData)
		if err != nil {
			base.LogFatal("LoadFile error [conf.json]: ", err)
		}
	}
	return configData
}
