package config

import (
	"encoding/json"
	"io/ioutil"
)

type UserPageConfig struct {
	TriggerRegex  string `json:"triggerregex"`
	NameRegex     string `json:"nameregex"`
	PositionRegex string `json:"positionregex"`
}

type CrawlConfig struct {
	Onion    string         `json:"onion"`
	Base     string         `json:"base"`
	Exclude  []string       `json:"exclude"`
	UserPage UserPageConfig `json:"userpage"`
}

func LoadCrawlConfig(filename string) (CrawlConfig, error) {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		return CrawlConfig{}, err
	}
	res := CrawlConfig{}
	err = json.Unmarshal(dat, &res)
	return res, err
}
