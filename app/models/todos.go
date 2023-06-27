package models

import (
	"log"
	"time"
)

type Todo struct {
	ID        int
	Content   string
	UserID    int
	CreatedAt time.Time
}

func (u *User) CreateTodo(content string) (err error) {
	cmd := `insert into todos (
		content,
		user_id,
		created_at
		) values (?,?,?)`

	_, err = Db.Exec(cmd, content, u.ID, time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetTodo(id int) (todo Todo, err error) {
	cmd := `select id, content, user_id, created_at from todos
	where id = ?`
	todo = Todo{}

	err = Db.QueryRow(cmd, id).Scan(
		&todo.ID,
		&todo.Content,
		&todo.UserID,
		&todo.CreatedAt)

	return todo, err
}

func (t *Todo) UpdateTodo(content string) (err error) {
	cmd := `update todos set content = ? where id = ?`
	_, err = Db.Exec(cmd, content, t.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
