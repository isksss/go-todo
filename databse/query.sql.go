// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: query.sql

package databse

import (
	"context"
)

const addTodo = `-- name: AddTodo :one
INSERT INTO todos (name) values(?) RETURNING id, name, completed
`

func (q *Queries) AddTodo(ctx context.Context, name string) (Todo, error) {
	row := q.db.QueryRowContext(ctx, addTodo, name)
	var i Todo
	err := row.Scan(&i.ID, &i.Name, &i.Completed)
	return i, err
}

const completedTodo = `-- name: CompletedTodo :one
UPDATE todos SET completed = '1' WHERE id = ? RETURNING id, name, completed
`

func (q *Queries) CompletedTodo(ctx context.Context, id int64) (Todo, error) {
	row := q.db.QueryRowContext(ctx, completedTodo, id)
	var i Todo
	err := row.Scan(&i.ID, &i.Name, &i.Completed)
	return i, err
}

const deleteTodo = `-- name: DeleteTodo :exec
DELETE FROM todos WHERE id = ?
`

func (q *Queries) DeleteTodo(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteTodo, id)
	return err
}

const getTodos = `-- name: GetTodos :many
SELECT id, name, completed FROM todos
`

func (q *Queries) GetTodos(ctx context.Context) ([]Todo, error) {
	rows, err := q.db.QueryContext(ctx, getTodos)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Todo
	for rows.Next() {
		var i Todo
		if err := rows.Scan(&i.ID, &i.Name, &i.Completed); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const unCompletedTodo = `-- name: UnCompletedTodo :one
UPDATE todos SET completed = '0' WHERE id = ? RETURNING id, name, completed
`

func (q *Queries) UnCompletedTodo(ctx context.Context, id int64) (Todo, error) {
	row := q.db.QueryRowContext(ctx, unCompletedTodo, id)
	var i Todo
	err := row.Scan(&i.ID, &i.Name, &i.Completed)
	return i, err
}
