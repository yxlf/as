package model

import "github.com/jinzhu/gorm"

type Store struct {
	*Model
	Name                 string  `json:"name" form:"name" binding:"required,max=20,min=6"`
	Threshold            float64 `json:"threshold" form:"threshold" binding:"required"`
	StartPercentagePoint uint8   `json:"start_percentage_point" form:"start_percentage_point" binding:"required,lte=100,gte=20"`
	IncrementValue       float64 `json:"increment_value" form:"increment_value" binding:"required,gt=0.0"`
	UpPercentagePoint    uint8   `json:"up_percentage_point" form:"up_percentage_point" binding:"required,lte=100"`
	CapPercentagePoint   uint8   `json:"cap_percentage_point" form:"cap_percentage_point" binding:"required,lte=100"`
}

func (s Store) TableName() string {
	return "stores"
}

func (s Store) Get(db *gorm.DB) (*Store, error) {
	err := db.Where(&s).First(&s).Error
	if err != nil {
		return nil, err
	}
	return &s, nil
}
func (s Store) Delete(db *gorm.DB) error {
	err := db.Select([]interface{}{"id"}).Delete(&s).Error
	if err != nil {
		return err
	}
	return nil
}
