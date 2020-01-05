package persist

import (
	"context"
	"log"

	"github.com/olivere/elastic"
)

func ItemServer() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0

		for {
			item := <-out
			log.Printf("Item serer: Got item #%d, %v", itemCount, item)
			itemCount++
			_, err := save(item)
			if err != nil {
				log.Print("Item saver: error saving item %v: %v", item, err)
			}
		}
	}()

	return out
}

func save(item interface{}) (string, error) {
	client, err := elastic.NewClient(
		// Must trun off sniff in docker
		elastic.SetSniff(false),
	)
	if err != nil {
		return "", err
	}

	resp, err := client.Index().
		Index("datiing_profile").
		Type("zhenai").
		BodyJson(item).
		Do(context.Background())
	if err != nil {
		return "", err
	}
	return resp.Id, nil
}
