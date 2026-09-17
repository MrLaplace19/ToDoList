package core_logger

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Lvl string `envconfig:"LEVEL" default:"DEBUG" required:"true"`
	Folder string `envconfig:"FOLDER" default:"./logs" required:"true"`
}

func NewConfig() (Config, error){
	var cfg Config

	if err :=envconfig.Process("LOGGER", &cfg); err != nil{
		return cfg, err
	}

	return cfg, nil
}

func NewConfigMust() Config{
	cfg, err := NewConfig()
	if err != nil{
		panic(err)
	}
	
	return cfg
}