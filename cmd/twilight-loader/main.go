package main

import (
	"encoding/json"
	"fmt"
	"os"

	twl "github.com/dendna/twilight-loader"
)

func main() {

	// var config = twl.Config{
	// 	Year: 2020,
	// 	TimezoneName:        "Europe/Moscow",
	// 	Latitude:            55.75,
	// 	Longitude:           37.8,
	// 	MorningTwilightType: twilight.DuskTypeCivil,
	// 	SunriseType:         twilight.DuskTypeSimple,
	// 	SunsetType:          twilight.DuskTypeSimple,
	// 	EveningTwilightType: twilight.DuskTypeCivil,
	// }

	config, err := loadConfiguration("config.json")
	if err != nil {
		fmt.Println("Error loading configuration : ", err)
		return
	}

	if err := twl.Generate(config); err != nil {
		fmt.Println("Error generating values: ", err)
	}
}

func loadConfiguration(file string) (twl.Config, error) {
	var config twl.Config
	configFile, err := os.Open(file)
	if err != nil {
		return twl.Config{}, err
	}
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	if err != nil {
		return twl.Config{}, err
	}
	return config, nil
}
