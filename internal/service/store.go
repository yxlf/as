package service

import (
	"as/internal/model"
)

func (s *Service) GetStoreByStoreName(storeName string) (*model.Store, error) {
	return s.dao.GetStoreByStoreName(storeName)

}

func (s *Service) CheckStoreNameIsUnique(name string) bool {
	storeName, err := s.GetStoreByStoreName(name)
	if err != nil && storeName == nil {
		return true
	}
	return false
}
