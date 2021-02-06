package database

import (
	"gorm.io/gorm"
)

//DB Create an exported global variable to hold database connection pool
var DB *gorm.DB
