package persist

import (
	"context"
	"fmt"
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
			save(item)
		}
	}()

	return out
}

func save(item interface{}) {
	client, err := elastic.NewClient(
		// Must trun off sniff in docker
		elastic.SetSniff(false),
	)
	if err != nil {
		panic(err)
	}

	resp, err := client.Index().
		Index("datiing_profile").
		Type("zhenai").
		BodyJson(item).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
