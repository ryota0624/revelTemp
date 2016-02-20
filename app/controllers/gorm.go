package controllers

import (
  "database/sql"
  _ "os"
  r "github.com/revel/revel"
  "github.com/jinzhu/gorm"
  _ "github.com/lib/pq"
  "MA/app/models"
)

type GormController struct {
  *r.Controller
  Txn *gorm.DB
}

var Gdb gorm.DB

func InitDB() {
  var err error
  Gdb, err = gorm.Open("postgres", "user=ryota dbname=golang sslmode=disable")
  if err != nil {
    r.ERROR.Println("FATAL", err)
    panic( err )
  }
  Gdb.AutoMigrate(&models.User{})
}

func (c *GormController) Begin() r.Result {
  txn := Gdb.Begin()
  if txn.Error != nil {
    panic(txn.Error)
  }
  c.Txn = txn
  return nil
}

func (c *GormController) Commit() r.Result {
  if c.Txn == nil {
    return nil
  }
  c.Txn.Commit()
  if err:= c.Txn.Error; err != nil && err != sql.ErrTxDone {
    panic(err)
  }
  c.Txn = nil
  return nil
}

func (c *GormController) Rollback() r.Result {
  if c.Txn == nil {
    return nil
  }
  c.Txn.Rollback()
  if err := c.Txn.Error; err != nil && err != sql.ErrTxDone {
    panic(err)
  }
  c.Txn = nil
  return nil
}