package alien

import (
	"log"

	modelBaak "github.com/kucinghitam13/alien/internal/model/baak"
)

func (a *Alien) GetJadkul(query string) (jadkulList modelBaak.JadkulList, err error) {
	jadkulList, err = a.repoBAAKWeb.GetJadkul(query)
	if err != nil {
		log.Println("[ERR] error fetching jadwal kuliah from repo : %s\n", err.Error())
		return
	}
	return
}
