package config

type Config struct {
	Env *Env
}

func SetConfigs(file string) (*Config, error) {
	//Loading Env Vars
	env, err := LoadEnvs(file)
	if err != nil {
		return &Config{}, err
	}

	c := &Config{
		Env: env,
	}
	return c, nil
}
