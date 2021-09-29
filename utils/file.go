package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

//确定文件是否存在

func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

//将结构写入yaml文件
func WriteToYaml(src string, conf interface{}) {
	data, err := yaml.Marshal(conf)
	if err != nil {
		log.Errorf("Conf转[]byte失败:%s", err.Error())
	}
	err = ioutil.WriteFile(src, data, 0777)
	if err != nil {
		log.Errorf("写yml文件失败", err.Error())
	}
}

// 嵌套创建文件
func CreatNestedFile(path string) (*os.File, error) {
	basePath := filepath.Dir(path)
	if !Exists(basePath) {
		err := os.MkdirAll(basePath, 0700)
		if err != nil {
			log.Errorf("无法创建目录，%s", err)
			return nil, err
		}
	}
	return os.Create(path)
}
