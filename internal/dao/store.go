package dao

import "as/internal/model"

type Store struct {
}

func (dao *Dao) GetStoreByStoreName(storeName string) (*model.Store, error) {
	m := new(model.Store)
	m.Name = storeName
	return m.Get(dao.engine)
}
