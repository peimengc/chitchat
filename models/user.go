package models

import "time"

type User struct {
	Id        int
	Uuid      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

func (user *User) CreateSession() (session Session, err error) {

	stmtIn, err := Db.Prepare("insert into sessions (uuid, email, user_id, created_at) values (?, ?, ?, ?)")

	if err != nil {
		return
	}

	defer stmtIn.Close()

	uuid := CreateUUID()

	_, err = stmtIn.Exec(uuid, user.Email, user.Id, time.Now())

	if err != nil {
		return
	}

	stmtOut, err := Db.Prepare("select id, email, uuid, user_id ,created_at from sessions where uuid = ?")

	if err != nil {
		return
	}

	err = stmtOut.QueryRow(uuid).Scan(&session.Id, &session.Email, &session.Uuid, &session.UserId, &session.CreatedAt)

	return
}

func (user *User) Session() (session Session, err error) {
	session = Session{}
	err = Db.QueryRow("select id, user_id, uuid, email, created_at from sessions where user_id = ?", user.Id).Scan(&session.Id, &session.UserId, &session.Uuid, &session.Email, &session.CreatedAt)
	return
}

func (user *User) Create() (err error) {
	uuid := CreateUUID()

	_, err = Db.Exec("insert into users (uuid, name, email, password, created_at) values(?, ?, ?, ?, ?)", uuid, user.Name, user.Email, Encrypt(user.Password), time.Now())

	if err != nil {
		return
	}

	err = Db.QueryRow("select id, uuid, created_at from users where uuid = ?", uuid).Scan(&user.Id, &user.Uuid, &user.CreatedAt)

	return
}

func (user *User) Delete() (err error) {

	_, err = Db.Exec("delete from users where id = ?", user.Id)

	return
}

func (user *User) Update() (err error) {
	_, err = Db.Exec("update users set name = ?, email = ? where id = ?", user.Name, user.Email, user.Id)
	return
}

func (user *User) UserDeleteAll() (err error) {
	_, err = Db.Exec("delete from users")
	return
}

func (user *User) CreateThread(topic string) (thread Thread, err error) {
	uuid := CreateUUID()

	_, err = Db.Exec("insert into threads (uuid,topic,user_id,created_at) values(?,?,?,?)", uuid, topic, user.Id, time.Now())

	if err != nil {
		return
	}

	err = Db.QueryRow("select id,uuid,topic,user_id,created_at from threads where uuid = ?", uuid).Scan(&thread.Id, &thread.Uuid, &thread.Topic, &thread.UserId, &thread.CreatedAt)
	return
}

func (user *User) CreatePost(thread Thread, body string) (post Post, err error) {
	uuid := CreateUUID()

	_, err = Db.Exec("insert into posts (uuid,body,user_id,thread_id,created_at) values(?,?,?,?,?)", uuid, body, user.Id, thread.Id, time.Now())

	if err != nil {
		return
	}

	err = Db.QueryRow("select id,uuid,body,user_id,thread_id,created_at from posts where uuid = ?", uuid).Scan(&post.Id, &post.Uuid, &post.Uuid, &post.UserId, &post.ThreadId, &post.CreatedAt)

	return
}

func Users() (users []User, err error) {
	rows, err := Db.Query("select id,uuid,name,email,password,created_at from users")

	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		user := User{}
		if err = rows.Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt); err != nil {
			return
		}
		//追加
		users = append(users, user)
	}

	return
}

func UserByUuid(uuid string) (user User, err error) {
	err = Db.QueryRow("select id,uuid,name,email,password,created_at from users where uuid = ?", uuid).Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return
}

func UserByEmail(email string) (user User, err error) {
	err = Db.QueryRow("select id,uuid,name,email,password,created_at from users where email = ?", email).Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return
}
