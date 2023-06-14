package config

import (
	"github.com/opensourceways/server-common-lib/utils"

	"github.com/opensourceways/app-cla-signing/cla/domain"
	"github.com/opensourceways/app-cla-signing/cla/domain/dp"
	"github.com/opensourceways/app-cla-signing/cla/infrastructure/repositoryimpl"
	"github.com/opensourceways/app-cla-signing/common/infrastructure/mongodb"
)

func LoadConfig(path string) (*Config, error) {
	cfg := new(Config)
	if err := utils.LoadFromYaml(path, cfg); err != nil {
		return nil, err
	}

	cfg.SetDefault()
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}

type configValidate interface {
	Validate() error
}

type configSetDefault interface {
	SetDefault()
}

type domainConfig struct {
	domain.Config

	DomainPrimitive dp.Config `json:"domain_primitive"  required:"true"`
}

type mongodbConfig struct {
	DB mongodb.Config `json:"db" required:"true"`

	repositoryimpl.Config
}

type Config struct {
	Mongodb mongodbConfig `json:"mongodb"     required:"true"`
	CLA     domainConfig  `json:"cla"         required:"true"`
}

func (cfg *Config) configItems() []interface{} {
	return []interface{}{
		&cfg.Mongodb.DB,
		&cfg.Mongodb.Config,
		&cfg.CLA.Config,
		&cfg.CLA.DomainPrimitive,
	}
}

func (cfg *Config) SetDefault() {
	items := cfg.configItems()
	for _, i := range items {
		if f, ok := i.(configSetDefault); ok {
			f.SetDefault()
		}
	}
}

func (cfg *Config) Validate() error {
	if _, err := utils.BuildRequestBody(cfg, ""); err != nil {
		return err
	}

	items := cfg.configItems()
	for _, i := range items {
		if f, ok := i.(configValidate); ok {
			if err := f.Validate(); err != nil {
				return err
			}
		}
	}

	return nil
}
