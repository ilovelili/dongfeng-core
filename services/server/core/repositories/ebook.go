package repositories

import (
	"github.com/ilovelili/dongfeng-core/services/server/core/models"
)

// EbookRepository ebook repository
type EbookRepository struct{}

// NewEbookRepository init ebook repository
func NewEbookRepository() *EbookRepository {
	return &EbookRepository{}
}

// Upsert upsert ebook
func (r *EbookRepository) Upsert(ebook *models.Ebook) (dirty bool, err error) {
	query := Table("ebooks").Alias("e").
		Project(
			"e.id",
			"e.hash",
		).
		Where().
		Eq("e.year", ebook.Year).
		Eq("e.class", ebook.Class).
		Eq("e.name", ebook.Name).
		Eq("e.date", ebook.Date).
		Sql()

	var _ebook models.Ebook
	err = session().Find(query, nil).Single(&_ebook)
	if err != nil || 0 == _ebook.ID {
		dirty = true
		err = session().Insert(ebook)
	} else if _ebook.Hash != ebook.Hash /*if hash same, do not update*/ {
		dirty = true
		ebook.ID = _ebook.ID
		ebook.Converted = false
		err = session().Update(ebook)
	} else {
		dirty = false
	}

	return
}
