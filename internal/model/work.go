package model

import "time"

type Work struct {
	*Model
	StoreId                 uint64    `json:"store_id" form:"store_id" binding:"required,gt=0"`
	Name                    string    `json:"name" form:"name" binding:"required,min=2,max=20"`
	Sex                     uint8     `json:"sex" form:"sex" binding:"required,oneof=0 1"`
	JoinDatetime            time.Time `json:"join_datetime" form:"join_datetime" binding:"required"`
	CanBonesetting          uint8     `json:"can_bonesetting" form:"can_bonesetting" binding:"required,oneof=0 1"`
	IsGraduationFromCollege uint8     `json:"is_graduation_from_college" form:"is_graduation_from_college" binding:"required,oneof=0 1"`
	IsProfessionalDocker    uint8     `json:"is_professional_docker" form:"is_professional_docker" binding:"required,oneof=0 1"`
	IsAssist                uint8     `json:"is_assist" form:"is_assist" binding:"required,oneof=0 1"`
}

func (w Work) TableName() string {
	return "workers"
}
