package models

import (
	"time"
)

type Thread struct {
	Id        int
	Uuid      string
	Topic     string
	UserId    int
	CreatedAt time.Time
}

func (t *Thread) CreatedAtDate() string {
	return t.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

func (t *Thread) NumReplies() (count int) {
	Db.QueryRow("select count(*) from posts where thread_id = ?", t.Id).Scan(&count)
	return
}

func Threads() (threads []Thread, err error) {
	rows, err := Db.Query("select id,uuid,topic,user_id,created_at from threads")

	for rows.Next() {
		thread := Thread{}
		if err = rows.Scan(&thread.Id, &thread.Uuid, &thread.Topic, &thread.UserId, &thread.CreatedAt); err != nil {
			return
		}
		threads = append(threads, thread)
	}

	rows.Close()
	return
}

func ThreadByUuid(uuid string) (thread Thread, err error) {
	err = Db.QueryRow("select id,uuid,topic,user_id,created_at from threads where uuid = ?", uuid).Scan(&thread.Id, &thread.Uuid, &thread.Topic, &thread.UserId, &thread.CreatedAt)
	return
}

func (t *Thread) User() (user User, err error) {
	err = Db.QueryRow("select id,uuid,name,email,password,created_at from users where id = ?", t.UserId).Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return
}

func (t *Thread) Posts() (posts []Post, err error) {
	rows, err := Db.Query("select id, uuid, body, user_id, thread_id, created_at from posts where thread_id = ?", t.Id)
	if err != nil {
		return
	}
	for rows.Next() {
		post := Post{}
		if err = rows.Scan(&post.Id, &post.Uuid, &post.Body, &post.UserId, &post.ThreadId, &post.CreatedAt); err != nil {
			return
		}
		posts = append(posts, post)
	}
	rows.Close()
	return
}
