package models

import "time"

type Post struct {
	Id        int
	Uuid      string
	Body      string
	UserId    int
	ThreadId  int
	CreatedAt time.Time
}

func (p *Post) CreatedAtDate() string {
	return p.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

func (p *Post) User() (user User, err error) {
	err = Db.QueryRow("select id,uuid,name,email,password,created_at from users where id = ?", p.UserId).Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return
}
