package model


import "gorm.io/datatypes"


type Flow struct {
ID string `gorm:"primaryKey"`
Name string
Nodes datatypes.JSON `gorm:"type:jsonb"`
}