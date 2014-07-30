package controllers

import (
  "App42PaaS-Revel-PostgreSQL-Sample/app/models"
  "github.com/revel/revel"
  "fmt"
)

type App struct {
	ModelController
}

func (c App) New() revel.Result {
	return c.Render()
}

func (c App) Index(name string, email string, description string) revel.Result{
  user := models.User{Name: name, Email: email, Description: description}
  c.Orm.Save(&user)

  rows, _ := c.Orm.Model(models.User{}).Rows()
  fmt.Println("All Data:", rows)
  
  users := []models.User{}
	for rows.Next() {
		user := models.User{}
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Description)
		fmt.Println(err)
		fmt.Println(user)
		users = append(users, user)
	}
	fmt.Println("Users:", users)

	return c.Render(users)
}
