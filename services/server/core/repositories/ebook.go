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

// Select select ebooks
func (r *EbookRepository) Select(year, class, name string) (ebooks []*models.Ebook, err error) {
	querybuilder := Table("ebooks").Alias("e").Where()

	if year != "" {
		querybuilder.Eq("e.year", year)
	}
	if class != "" {
		querybuilder.Eq("e.class", class)
	}
	if name != "" {
		querybuilder.Eq("e.name", name)
	}

	query := querybuilder.Eq("e.converted", 1).Sql()
	// no rows is actually not an error
	if err = session().Find(query, nil).All(&ebooks); err != nil && norows(err) {
		err = nil
	}

	return
}

// Upsert upsert ebook
func (r *EbookRepository) Upsert(ebook *models.Ebook, force bool) (dirty bool, err error) {
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
		err = session().Update(ebook)
	} else {
		dirty = false
		// force update
		if force {
			ebook.ID = _ebook.ID
			err = session().Update(ebook)
		}
	}

	return
}
