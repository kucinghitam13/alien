package web

import (
	"errors"
	"log"
	"net/http"
)

type Repo struct {
	Config      *Config
	Credentials *Credentials
	HTTPClient  *http.Client
}

type Config struct {
	Host                 string
	JadkulEndpoint       string
	RetryMaxAttempt      int
	RetryIntervalSeconds int
}

type Credentials struct {
	Cookies   []*http.Cookie
	CSRFToken string
}

func New(config *Config) (repo *Repo, err error) {
	if config == nil {
		err = errors.New("Config is nil")
		return
	}

	repo = &Repo{
		Config:      config,
		Credentials: &Credentials{},
		HTTPClient:  &http.Client{},
	}
	if err = repo.refreshCredentials(); err != nil {
		log.Printf("[ERR] error initializing credentials (max retry attempted): %s\n", err.Error())
		return
	}
	return
}
