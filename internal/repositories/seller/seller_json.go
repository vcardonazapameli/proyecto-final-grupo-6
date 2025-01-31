package seller

import (
	"github.com/arieleon_meli/proyecto-final-grupo-6/internal/utils/errors"
	"github.com/arieleon_meli/proyecto-final-grupo-6/pkg/models"
)

type SellerRepositoryJSON struct {
	db map[int]models.Seller
}

func NewSellerRepositoryJSON(db map[int]models.Seller) *SellerRepositoryJSON {
	// default db
	defaultDb := make(map[int]models.Seller)
	if db != nil {
		defaultDb = db
	}
	return &SellerRepositoryJSON{db: defaultDb}
}

func (r *SellerRepositoryJSON) GetAll() (s map[int]models.Seller, err error) {
	return r.db, err
}

func (r *SellerRepositoryJSON) GetByID(id int) (models.Seller, error) {
	if s, ok := r.db[id]; !ok {
		return models.Seller{}, errors.ErrorNotFound
	} else {
		return s, nil
	}
}

func (r *SellerRepositoryJSON) SearchByCID(cid int) (models.Seller, bool) {
	for _, s := range r.db {
		if s.Cid == cid {
			return s, true
		}
	}
	return models.Seller{}, false
}

func (r *SellerRepositoryJSON) GetBiggestID() (max int) {
	for _, s := range r.db {
		if s.Id > max {
			max = s.Id
		}
	}
	return
}

func (r *SellerRepositoryJSON) Save(s models.Seller) (models.Seller, error) {
	if _, exists := r.SearchByCID(s.Cid); exists {
		return models.Seller{}, errors.ErrorConflict
	}

	// If does not exist, save.
	s.Id = r.GetBiggestID() + 1

	r.db[s.Id] = s
	return s, nil
}

func (r *SellerRepositoryJSON) Delete(id int) error {
	if _, exists := r.db[id]; !exists {
		return errors.ErrorNotFound
	}
	delete(r.db, id)
	return nil
}

func (r *SellerRepositoryJSON) Update(s models.Seller) {
	r.db[s.Id] = s
}
