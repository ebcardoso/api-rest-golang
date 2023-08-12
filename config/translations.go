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
	Auth struct {
		Signup struct {
			Errors struct {
				AlreadyExists     string `yaml:"already_exists"`
				PasswordDifferent string `yaml:"password_different"`
				SaveUser          string `yaml:"save_user"`
			} `yaml:"errors"`
		} `yaml:"signup"`
	} `yaml:"auth"`
	Users struct {
		Errors struct {
			NotFound string `yaml:"not_found"`
		} `yaml:"errors"`
		Create struct {
			Errors string `yaml:"errors"`
		} `yaml:"create"`
		Load struct {
			Errors string `yaml:"errors"`
		} `yaml:"load"`
	} `yaml:"users"`
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
