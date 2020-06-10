package conf

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/elastic/go-elasticsearch/v7"
	"gopkg.in/yaml.v2"

	"go-web/web/library/igorm"
)

type Config struct {
	Mysql *igorm.MysqlConf `json:"mysql"`
	Elastic elasticsearch.Config `json:"elastic"`
}

var confPath string

func init() {
	flag.StringVar(&confPath, "c", "./config.yaml", "config path")
}

func Init()  {
	load()
}

var config = Config{}

func load() {
	file, err := ioutil.ReadFile(confPath)
	if err != nil {
		log.Fatal("config.yaml read err:", err)
	}
	if err := yaml.Unmarshal(file, &config); err != nil {
		log.Fatal("config Unmarshal err:", err)
	}
	fmt.Printf("%+v \n", &config)
}

func New() *Config {
	return &config
}
