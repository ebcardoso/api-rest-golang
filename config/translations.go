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
			Success     string `yaml:"success"`
			Invalid     string `yaml:"invalid"`
			UserBlocked string `yaml:"user_blocked"`
		} `yaml:"signin"`
		ForgotPasswordToken struct {
			Success string `yaml:"success"`
			Errors  struct {
				Default string `yaml:"default"`
			} `yaml:"errors"`
		} `yaml:"forgot_password_token"`
	} `yaml:"auth"`
	Users struct {
		Errors struct {
			NotFound string `yaml:"not_found"`
		} `yaml:"errors"`
		Create struct {
			Errors string `yaml:"errors"`
		} `yaml:"create"`
		List struct {
			Errors string `yaml:"errors"`
		} `yaml:"list"`
		Fetch struct {
			Errors string `yaml:"errors"`
		} `yaml:"fetch"`
		Load struct {
			Errors string `yaml:"errors"`
		} `yaml:"load"`
		Update struct {
			Errors string `yaml:"errors"`
		} `yaml:"update"`
		Destroy struct {
			Success string `yaml:"success"`
			Errors  string `yaml:"errors"`
		}
		Block struct {
			Success string `yaml:"success"`
			Errors  string `yaml:"errors"`
		} `yaml:"block"`
		Unblock struct {
			Success string `yaml:"success"`
			Errors  string `yaml:"errors"`
		} `yaml:"unblock"`
	} `yaml:"users"`
	Errors struct {
		InvalidParams string `yaml:"invalid_params"`
		ParseId       string `yaml:"parse_id"`
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
