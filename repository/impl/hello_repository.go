package impl

import (
	"filesystem/entity"
	"fmt"
	"github.com/pkg/errors"
	"xorm.io/xorm"
)

type HelloRepository struct {
	*xorm.Session
}

func (h *HelloRepository) Get() *entity.User {
	u := &entity.User{}
	h.Session.Table("user").Where("id=?", "1").Get(u)
	return u
}

func (h *HelloRepository) Set(user *entity.User) error {
	is, err := h.Session.Table("user").InsertOne(user)
	fmt.Println(is)
	return errors.WithStack(err)
}
