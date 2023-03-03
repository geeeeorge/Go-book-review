package model

import "github.com/geeeeorge/Go-book-review/src/app/dao"

type User struct {
	ID       int64
	Username string
	Password string
}

func (u *User) LoadDAO(d *dao.User) {
	u.ID = d.ID
	u.Username = d.Username
	u.Password = d.Password
}

func (u *User) DAO() *dao.User {
	ret := &dao.User{
		ID:       u.ID,
		Username: u.Username,
		Password: u.Password,
	}
	return ret
}
