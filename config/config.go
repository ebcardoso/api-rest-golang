package config

import "go.mongodb.org/mongo-driver/mongo"

type Config struct {
	Env          *Env
	Database     *mongo.Database
	Translations *Translations
}

func SetConfigs(file string) (*Config, error) {
	//Loading Env Vars
	env, err := LoadEnvs(file)
	if err != nil {
		return &Config{}, err
	}

	//Loading Database
	db, err := LoadDatabase(env.MONGO_URI, env.MONGO_DATABASE)
	if err != nil {
		return &Config{}, err
	}

	//Loading Translations
	translations, err := LoadTranslations(env.DEFAULT_TRANSLATION)
	if err != nil {
		return &Config{}, err
	}

	c := &Config{
		Env:          env,
		Database:     db,
		Translations: translations,
	}
	return c, nil
}
