package model

type User struct {
	Id     string `db:"id"`
	Name   string `db:"name"`
	Gender string `db:"gender"`
}

func (User) TableName() string {
	return "user"
}
