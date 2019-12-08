package alien

import (
	"errors"

	repoBAAK "github.com/kucinghitam13/alien/internal/repository/baak"
)

type Alien struct {
	Config      *Config
	repoBAAKWeb repoBAAK.Repo
}

type Config struct {
}

func New(config *Config, repoBAAKWeb repoBAAK.Repo) (usecase *Alien, err error) {
	if repoBAAKWeb == nil {
		err = errors.New("repo BAAK web is nil")
		return
	}

	usecase = &Alien{
		Config:      config,
		repoBAAKWeb: repoBAAKWeb,
	}

	return
}
