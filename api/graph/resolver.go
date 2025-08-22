package graph


import "gorm.io/gorm"


// Resolver держит зависимости, доступные в резолверах
// (БД, кэш, конфиг и т.п.)
type Resolver struct { DB *gorm.DB }