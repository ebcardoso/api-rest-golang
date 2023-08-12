package config

import (
	"errors"
	"os"

	"gopkg.in/yaml.v2"
)

var (
	ErrLoadTranslation  = "error to load .yml file"
	ErrParseTranslation = "error to parse .yml file"
)

type Translations struct {
}

func LoadTranslations(fileYml string) (*Translations, error) {
	file, err := os.Open("config/translations/" + fileYml + ".yml")
	if err != nil {
		return &Translations{}, errors.New(ErrLoadTranslation)
	}

	defer file.Close()

	translations := Translations{}
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&translations)
	if err != nil {
		return &Translations{}, errors.New(ErrParseTranslation)
	}

	return &translations, nil
}
