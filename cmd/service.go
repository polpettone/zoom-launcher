package cmd

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Zooms struct {
	Entries map[string][]string
}

func LoadZooms(file string) (*Zooms, error) {
	content, err := loadRawZooms(file)
	if err != nil {
		return nil, err
	}
	return convertRawZoomsToZooms(content)
}

func convertRawZoomsToZooms(content []byte) (*Zooms, error) {
	zooms := &Zooms{}
	err := yaml.Unmarshal(content, &zooms)
	if err != nil {
		return nil, err
	}
	return zooms, nil
}

func loadRawZooms(file string) ([]byte, error) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return content, nil
}
