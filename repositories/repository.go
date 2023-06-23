package repositories

import (
	"game-boost/db"

	"gorm.io/gorm"
)

var DB *gorm.DB = db.Init()
