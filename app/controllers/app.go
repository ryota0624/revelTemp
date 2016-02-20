package controllers

import (
  "github.com/revel/revel"
  "MA/app/models"
)

type App struct {
	GormController
}

func (c App) Index() revel.Result {
  users := [] models.User{}
  usernames := [] string {}
  c.Txn.Find(&users)
  for _, user := range users {
    usernames = append(usernames, user.Name)
  }
  // greeting := "OK"
	return c.Render(usernames)
}

func (c App) Hello(myName string) revel.Result {
  c.Validation.Required(myName).Message("your name is required")
  c.Validation.MinSize(myName, 3).Message("your name is not long enogh")
  
  if c.Validation.HasErrors() {
    c.Validation.Keep()
    c.FlashParams()
    return c.Redirect(App.Index)
  }
  return c.Render(myName)
}

func (c App) Create(username string) revel.Result {
  c.Validation.Required(username).Message("your name is required")
  c.Validation.MinSize(username, 3).Message("your name is not long enogh")
  if c.Validation.HasErrors() {
    c.Validation.Keep()
    c.FlashParams()
    return c.Redirect(App.Index)
  }
  user := models.User{Name: username}
  c.Txn.Create(&user)
	return c.RenderJson(user)
}
