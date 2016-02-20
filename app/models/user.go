package models

import (
  _ "database/sql"
  _ "github.com/go-sql-driver/mysql"
  "github.com/jinzhu/gorm"
  _ "github.com/lib/pq"
)

type User struct {
  gorm.Model
  Name string `sql:"default:'unko'"`
}
