package animapi

import "github.com/otiai10/animapi/model"
import "github.com/otiai10/animapi/infrastructure"

func (m *MySQL) FindAnisongsByTID(tid int) (anisongs []model.Anisong) {
	table := infrastructure.NewAnisongsTable(m.db)
	if e := table.CreateIfNotExists(); e != nil {
		return
	}
	rows, e := table.FindByTID(tid)
	if e != nil {
		return
	}
	return model.CreateAnisongsFromMySQLResponse(rows)
}
func (m *MySQL) AddAnisongsOfAnime(anime model.Anime) (e error) {
	table := infrastructure.NewAnisongsTable(m.db)
	if e = table.CreateIfNotExists(); e != nil {
		return
	}
	for _, ansng := range anime.Anisongs {
		if e = table.Add(ansng.TID, ansng.Title, ansng.Label, ansng.Index, ansng.Detail); e != nil {
			return
		}
	}
	return
}
func (m *MySQL) DeleteAnisongOfAnime(anime model.Anime) (e error) {
	table := infrastructure.NewAnisongsTable(m.db)
	if e = table.CreateIfNotExists(); e != nil {
		return
	}
	return table.DeleteByTID(anime.TID)
}
