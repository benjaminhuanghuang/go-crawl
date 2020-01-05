package persist

import (
	"context"
	"encoding/json"
	"testing"

	"../engine"
	"../model"
	"github.com/olivere/elastic"
)

func TestSave(t *testing.T) {
	expected := engine.Item{
		Url:  "http://album.zhenai.com/u/108906739",
		Type: "zhenai",
		Id:   "108906739",
		Payload: model.Profile{
			Name:       "某某某",
			Gender:     "男",
			Age:        33,
			Height:     111,
			Weight:     111,
			Income:     "3330-5000",
			Marriage:   "离异",
			Education:  "大学",
			Occupation: "大学",
			Hokou:      "上海",
			Xinzuo:     "天平",
			House:      "有房",
			Car:        "无车",
		},
	}

	err := save(expected)

	if err != nil {
		panic(err)
	}

	client, err := elastic.NewClient(
		// Must trun off sniff in docker
		elastic.SetSniff(false),
	)
	if err != nil {
		panic(err)
	}

	resp, err := client.Get().
		Index("datiing_profile").
		Type(expected.Type).
		Id(expected.Id).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	var actual engine.Item
	err = json.Unmarshal([]byte(resp.Source), &actual)
	if err != nil {
		panic(err)
	}

	actualProfile, err := model.FromJsonObj(actual.Payload)

	actual.Payload = actualProfile

	if actual != expected {
		t.Errorf("got %v; expected %v", actual, expected)
	}
}
