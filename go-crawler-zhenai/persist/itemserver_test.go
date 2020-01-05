package persist

import (
	"testing"
	"../model"
)

func TestSave(t *testing.T) {
	profile := model.Profile{
		Name   :    "某某某",
		Gender :    "男",
		Age :       33,
		Height :    111,
		Weight :    111,
		Income :    "3330-5000",
		Marriage:   "离异",
		Education:  "大学",
		Occupation: "大学",
		Hokou :     "上海",
		Xinzuo:     "天平",
		House :     "有房",
		Car :       "无车",
	}

	save(profile)

}
