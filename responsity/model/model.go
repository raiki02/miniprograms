package model

import (
	"gorm.io/gorm"
)

type MiniPrograms struct {
	Name    string `gorm:"unique"`
	Status  bool
	Version string
	gorm.Model
}
