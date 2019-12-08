package baak

import "github.com/kucinghitam13/alien/internal/model/baak"

type Repo interface {
	GetJadkul(query string) (jadkulList baak.JadkulList, err error)
}
