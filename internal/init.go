package app

import (
	"github.com/julienschmidt/httprouter"
	"github.com/kucinghitam13/alien/internal/config"
	delivery "github.com/kucinghitam13/alien/internal/delivery/alien/http"
	repoBaakWeb "github.com/kucinghitam13/alien/internal/repository/baak/web"
	"github.com/kucinghitam13/alien/internal/usecase"
	usecaseAlien "github.com/kucinghitam13/alien/internal/usecase/alien"
)

const (
	API_PREFIX = "/api"
)

type App struct {
	UsecaseAlien usecase.Alien
}

func Init(config *config.Config) (app *App, err error) {
	repoBaakWeb, err := repoBaakWeb.New(&repoBaakWeb.Config{
		Host:                 config.BAAK.Host,
		JadkulEndpoint:       config.BAAK.JadkulEndpoint,
		RetryMaxAttempt:      config.BAAK.RetryMaxAttempt,
		RetryIntervalSeconds: config.BAAK.RetryIntervalSeconds,
	})
	if err != nil {
		return
	}

	configAlien := &usecaseAlien.Config{}

	usecaseAlien, err := usecaseAlien.New(configAlien, repoBaakWeb)
	if err != nil {
		return
	}

	app = &App{
		UsecaseAlien: usecaseAlien,
	}

	return
}

func (a *App) SetEndpoint(router *httprouter.Router) (err error) {
	handler, err := delivery.New(a.UsecaseAlien)
	if err != nil {
		return
	}

	router.GET(API_PREFIX+"/jadwal-kuliah", handler.GetJadkul)

	return
}
