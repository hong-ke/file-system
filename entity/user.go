package entity

type User struct {
	ID       string `xorm:"id"`
	Name     string
	Password string
}
