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
		Protector struct {
			NotAuth       string `yaml:"not_auth"`
			TokenRequired string `yaml:"token_required"`
		} `yaml:"protector"`
		Signup struct {
			Errors struct {
				AlreadyExists     string `yaml:"already_exists"`
				PasswordDifferent string `yaml:"password_different"`
				SaveUser          string `yaml:"save_user"`
			} `yaml:"errors"`
		} `yaml:"signup"`
		Signin struct {
			Success string `yaml:"success"`
			Invalid string `yaml:"invalid"`
		} `yaml:"signin"`
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
		Destroy struct {
			Success string `yaml:"success"`
			Errors  string `yaml:"errors"`
		}
	} `yaml:"users"`
	Errors struct {
		ParseId string `yaml:"parse_id"`
	} `yaml:"errors"`
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
