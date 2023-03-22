-- name: GetTodos :many
SELECT * FROM todos;

-- name: AddTodo :one
INSERT INTO todos (name) values(?) RETURNING *;

-- name: CompletedTodo :one
UPDATE todos SET completed = '1' WHERE id = ? RETURNING *;

-- name: UnCompletedTodo :one
UPDATE todos SET completed = '0' WHERE id = ? RETURNING *;

-- name: DeleteTodo :exec
DELETE FROM todos WHERE id = ?;

-- -- name: CheckTable :one
-- SELECT COUNT(*) FROM sqlite_master WHERE TYPE='table' AND name=?;