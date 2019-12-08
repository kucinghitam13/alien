package usecase

import modelBaak "github.com/kucinghitam13/alien/internal/model/baak"

type Alien interface {
	GetJadkul(query string) (jadkulList modelBaak.JadkulList, err error)
}
