package http

import (
	"errors"
	"log"

	"github.com/kucinghitam13/alien/internal/usecase"
)

type Handler struct {
	Service usecase.Alien
}

func New(service usecase.Alien) (handler *Handler, err error) {
	if service == nil {
		err = errors.New("service is nil")
		log.Printf("[ERR] error initiating handler : %s\n", err)
		return
	}

	handler = &Handler{
		Service: service,
	}
	return
}
