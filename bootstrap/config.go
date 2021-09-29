package bootstrap

import (
	"io/ioutil"
	"steamBackend/conf"
	"steamBackend/utils"
	"strings"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// 读取配置文件

func ReadConf(config string) bool {
	log.Infof("读取配置文件%s...", config)
	if !utils.Exists(config) {
		log.Infof("找不到配置文件:%s", config)
		if !Write(config) {
			return false
		}
	}
	confFile, err := ioutil.ReadFile(config)
	if err != nil {
		log.Errorf("读取配置文件时发生错误:%s", err.Error())
		return false
	}
	err = yaml.Unmarshal(confFile, conf.Conf)
	if err != nil {
		log.Errorf("加载配置文件时发生错误:%s", err.Error())
		return false
	}
	log.Debugf("config:%+v", conf.Conf)
	conf.Origins = strings.Split(conf.Conf.Server.SiteUrl, ",")
	return true
}

func Write(path string) bool {
	log.Infof("创建默认配置文件")
	file, err := utils.CreatNestedFile(path)
	if err != nil {
		log.Errorf("无法创建配置文件, %s", err)
		return false
	}
	defer func() {
		_ = file.Close()
	}()
	str := `
server:
  name: "SteamBackend"
  server_name: "俄罗斯圣彼得堡"
  address: "0.0.0.0"
  port: "5244"
  site_url: "*"
`
	_, err = file.WriteString(str)
	if err != nil {
		log.Errorf("无法写入配置文件, %s", err)
		return false
	}
	return true
}
