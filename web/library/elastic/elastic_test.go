package elastic_test

import (
	"flag"
	"log"
	"testing"

	"go-web/web/conf"
	"go-web/web/library/elastic"
)


func TestNew(t *testing.T) {
	flag.Set("c", "../../config.yaml")
	flag.Parse()
	conf.Init()
	client:=elastic.New(conf.New().Elastic)
	if _, err := client.Ping();err!=nil{
		log.Fatal(err)
	}
}
