package models

import "encoding/gob"

type User struct {
	ID         NullInt32  `json:"id" form:"id" db:"id"`
	Name       NullString `json:"name" form:"name" db:"name"`
	Email      NullString `json:"email" form:"email" db:"email"`
	Mobile     NullString `json:"mobile" form:"mobile" db:"mobile"`
	Password   NullString `json:"password" form:"password" db:"password"`
	Lock       NullBool   `json:"lock" form:"lock" db:"lock"`
	Salt       NullString `db:"salt"`
	CreateTime NullInt64  `db:"create_time"`
	UpdateTime NullInt64  `db:"update_time"`
	LoginTime  NullInt64  `db:"login_time"`
}

func init() {
	gob.Register(User{})
}
