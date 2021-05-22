package model

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
