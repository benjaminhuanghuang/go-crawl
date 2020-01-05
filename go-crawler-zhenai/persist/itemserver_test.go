package persist

import (
	"context"
	"encoding/json"
	"testing"

	"../model"
	"github.com/olivere/elastic"
)

func TestSave(t *testing.T) {
	profile := model.Profile{
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
	}

	id, err := save(profile)

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
		Type("zhenai").
		Id(id).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	var actual model.Profile
	err = json.Unmarshal([]byte(resp.Source), &actual)
	if err != nil {
		panic(err)
	}
	if actual != profile {
		t.Errorf("got %v; expected %v", actual, profile)
	}
}
