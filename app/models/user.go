package models

type User struct {
  Id            int
  Name          string `sql:"type:varchar(100)"`
  Email         string `sql:"type:varchar(100)"`
  Description   string `sql:"type:varchar(100)"`
}
