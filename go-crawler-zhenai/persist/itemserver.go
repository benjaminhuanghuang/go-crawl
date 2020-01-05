package persist

import (
	"context"
	"errors"
	"log"

	"../engine"
	"github.com/olivere/elastic"
)

func ItemServer() chan engine.Item {
	out := make(chan engine.Item)
	go func() {
		itemCount := 0

		for {
			item := <-out
			log.Printf("Item serer: Got item #%d, %v", itemCount, item)
			itemCount++
			err := save(item)
			if err != nil {
				log.Printf("Item saver: error saving item %v: %v", item, err)
			}
		}
	}()

	return out
}

func save(item engine.Item) error {
	client, err := elastic.NewClient(
		// Must trun off sniff in docker
		elastic.SetSniff(false),
	)
	if err != nil {
		return err
	}

	if item.Type == "" {
		return errors.New("Must supply type")
	}
	indexService := client.Index().
		Index("datiing_profile").
		Type(item.Type).
		BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err = indexService.
		Do(context.Background())

	if err != nil {
		return err
	}
	return nil
}
