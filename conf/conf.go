package conf

import (
	"errors"
	"sync"

	"github.com/pelletier/go-toml"
)

type Service struct {
	Name    string `toml:"name"`
	Addr    string `toml:"addr"`
	Network string `toml:"network"`
}

var (
	ServiceConf *Service
	once        = new(sync.Once)

	ErrServiceConfigNotFound = errors.New("service config not found")
)

func ParseConfigFromFile(path string) error {
	c, err := toml.LoadFile(path)
	if err != nil {
		return err
	}

	once.Do(func() {
		if ServiceConf == nil {
			ServiceConf = new(Service)
		}
	})

	sc, ok := c.Get("service").(*toml.Tree)
	if !ok {
		return ErrServiceConfigNotFound
	}
	return sc.Unmarshal(ServiceConf)
}
