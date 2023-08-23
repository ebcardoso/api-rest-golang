package config

import (
	"errors"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

var (
	ErrLoadEnvs  = "error to load .env file"
	ErrParseEnvs = "error to parse .env file"
)

type Env struct {
	DEFAULT_TRANSLATION string `env:"DEFAULT_TRANSLATION,required"`
	JWT_KEY             string `env:"JWT_KEY,required"`
	MONGO_URI           string `env:"MONGO_URI,required"`
	MONGO_DATABASE      string `env:"MONGO_DATABASE,required"`
	SMTP_ADDRESS        string `env:"SMTP_ADDRESS,required"`
	SMTP_PORT           int    `env:"SMTP_PORT,required"`
	SMTP_SENDER_NAME    string `env:"SMTP_SENDER_NAME,required"`
	SMTP_USERNAME       string `env:"SMTP_USERNAME,required"`
	SMTP_PASSWORD       string `env:"SMTP_PASSWORD,required"`
}

func LoadEnvs(file string) (*Env, error) {
	err := godotenv.Load(file)
	if err != nil {
		return &Env{}, errors.New(ErrLoadEnvs)
	}

	envs := Env{}
	err = env.Parse(&envs)
	if err != nil {
		return &Env{}, errors.New(ErrParseEnvs)
	}
	return &envs, nil
}
