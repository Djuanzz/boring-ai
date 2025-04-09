package config

import "googlemaps.github.io/maps"

func NewMapsClient(apiKey string) *maps.Client {
	cli, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		panic("Error creating Google Maps client")
	}
	return cli
}
