package models

import "time"

type Session struct {
	Id        int
	Email     string
	Uuid      string
	UserId    int
	CreatedAt time.Time
}

func (session *Session) Check() (valid bool, err error) {
	err = Db.QueryRow("select id,email,uuid,user_id,created_at from sessions where uuid = ?", session.Uuid).Scan(&session.Id, &session.Email, &session.Uuid, &session.UserId, &session.CreatedAt)
	if err != nil {
		return
	}
	if session.Id != 0 {
		valid = true
	}
	return
}

func (session *Session) DeleteByUuid() (err error) {
	_, err = Db.Exec("delete from sessions where uuid = ?", session.Uuid)

	return
}

func (session *Session) User() (user User, err error) {
	err = Db.QueryRow("select id,uuid,name,email,password,created_at from users where id = ?", session.UserId).Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return
}

func SessionDeleteAll() (err error) {
	_, err = Db.Exec("delete from sessions")
	return
}
