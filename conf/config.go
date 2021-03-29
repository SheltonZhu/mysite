package conf

import (
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"log"
	"os"
	"path"
)

var Conf *BaseData

func init() {
	var err error
	cwd, _ := os.Getwd()
	Conf, err = ReadConf(path.Join(cwd, "conf", "config.conf"))
	if err != nil {
		log.Panicf("read config failed. %v", err)
	}
}

type BaseData struct {
	Db Databases `toml:"databases"`
}

type Databases struct {
	Mysql Database `toml:"mysql"`
}

type Database struct {
	ConnStr string `toml:"connStr"`
}

func ReadConf(fileName string) (p *BaseData, err error) {
	var (
		fp          *os.File
		fileContent []byte
	)
	p = new(BaseData)
	if fp, err = os.Open(fileName); err != nil {
		log.Panicln("open error ", err)
		return
	}

	if fileContent, err = ioutil.ReadAll(fp); err != nil {
		log.Panicln("ReadAll error ", err)
		return
	}

	if err = toml.Unmarshal(fileContent, p); err != nil {
		log.Panicln("toml.Unmarshal error ", err)
		return
	}
	return
}
