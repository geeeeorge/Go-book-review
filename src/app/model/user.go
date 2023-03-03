package model

import "github.com/geeeeorge/Go-book-review/src/app/dao"

type User struct {
	ID       int
	Username string
	Password string
}

func (u *User) LoadDAO(d *dao.User) {
	if d.ID != 0 {
		u.ID = d.ID
	}
	if d.Username != "" {
		u.Username = d.Username
	}
	if d.Password != "" {
		u.Password = d.Password
	}
}

func (u *User) DAO() *dao.User {
	ret := &dao.User{
		ID:       u.ID,
		Username: u.Username,
		Password: u.Password,
	}
	return ret
}
