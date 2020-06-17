package elastic_test

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"testing"

	"go-web/web/conf"
	"go-web/web/library/elastic"
)

func TestNew(t *testing.T) {
	flag.Set("c", "../../config.yaml")
	flag.Parse()
	conf.Init()
	client := elastic.New(conf.New().Elastic)
	if _, err := client.Ping(); err != nil {
		log.Fatal(err)
	}

	info, err := client.Info()
	if err != nil {
		log.Fatal(err)
	}
	if info.IsError() {
		log.Fatal("info err")
	}
	var msg map[string]interface{}
	if err := json.NewDecoder(info.Body).Decode(msg); err != nil {
		log.Fatal(err)
	}
	fmt.Println(msg)

}
