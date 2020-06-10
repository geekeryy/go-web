package elastic

import (
	"log"

	"github.com/elastic/go-elasticsearch/v7"
)

func New(c elasticsearch.Config) *elasticsearch.Client {
	client, err := elasticsearch.NewClient(c)
	if err != nil {
		log.Println("ES err:", err)
		panic(err)
	}
	return client
}

func CreateIndex(client *elasticsearch.Client)  {
	client.Index.WithFilterPath()
}



