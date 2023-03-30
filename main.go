package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// ToDo管理アプリケーション
// 標準ライブラリのnet/httpを使って作成
// 1. ToDoの追加
// 2. ToDoの一覧表示
// 3. ToDoの完了
// 4. ToDoの削除

// ToDoの構造体 jsonタグつき
type ToDo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// ToDoのスライス
var todos []ToDo

// ToDoのID
var id int

// メイン関数
func main() {
	// ルーティング
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/todos", todosHandler)
	http.HandleFunc("/todos/", todoHandler)

	// サーバー起動
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// ルートハンドラ
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World\n")
}

// ToDo一覧ハンドラ
func todosHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// ToDo一覧を取得
		getTodosHandler(w, r)
	case "POST":
		// ToDoを追加
		postTodosHandler(w, r)
	default:
		// 405 Method Not Allowed
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// ToDo一覧取得ハンドラ
func getTodosHandler(w http.ResponseWriter, r *http.Request) {
	// JSONを返す
	w.Header().Set("Content-Type", "application/json")
	// JSONエンコーダーを作成
	encoder := json.NewEncoder(w)
	// JSONエンコード
	encoder.Encode(todos)
}

// ToDo追加ハンドラ
func postTodosHandler(w http.ResponseWriter, r *http.Request) {
	// ToDoのタイトルを取得
	title := r.FormValue("title")
	// ToDoを追加
	todo := ToDo{
		ID:        id,
		Title:     title,
		Completed: false,
	}
	todos = append(todos, todo)
	id++
	// JSONを返す
	w.Header().Set("Content-Type", "application/json")
	// JSONエンコーダーを作成
	encoder := json.NewEncoder(w)
	// JSONエンコード
	encoder.Encode(todo)
}

// ToDoハンドラ
func todoHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// ToDoを取得
		getTodoHandler(w, r)
	case "PUT":
		// ToDoを更新
		putTodoHandler(w, r)
	case "DELETE":
		// ToDoを削除
		deleteTodoHandler(w, r)
	default:
		// 405 Method Not Allowed
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// ToDo取得ハンドラ
func getTodoHandler(w http.ResponseWriter, r *http.Request) {
	// ToDoのIDを取得
	id := r.URL.Path[len("/todos/"):]
	// ToDoを取得
	todo := findTodoByID(atoi(id))
	// JSONを返す
	w.Header().Set("Content-Type", "application/json")
	// JSONエンコーダーを作成
	encoder := json.NewEncoder(w)
	// JSONエンコード
	encoder.Encode(todo)
}

// ToDo更新ハンドラ
func putTodoHandler(w http.ResponseWriter, r *http.Request) {
	// ToDoのIDを取得
	id := r.URL.Path[len("/todos/"):]
	// ToDoを取得
	todo := findTodoByID(atoi(id))
	// ToDoのタイトルを更新
	todo.Title = r.FormValue("title")
	// ToDoの完了を更新
	todo.Completed = r.FormValue("completed") == "true"
	// JSONを返す
	w.Header().Set("Content-Type", "application/json")
	// JSONエンコーダーを作成
	encoder := json.NewEncoder(w)
	// JSONエンコード
	encoder.Encode(todo)
}

// ToDo削除ハンドラ
func deleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	// ToDoのIDを取得
	id := r.URL.Path[len("/todos/"):]
	// ToDoを取得
	todo := findTodoByID(atoi(id))
	// ToDoを削除
	deleteTodoByID(atoi(id))
	// JSONを返す
	w.Header().Set("Content-Type", "application/json")
	// JSONエンコーダーを作成
	encoder := json.NewEncoder(w)
	// JSONエンコード
	encoder.Encode(todo)
}

// ToDoをIDで検索
func findTodoByID(id int) *ToDo {
	for i, todo := range todos {
		if todo.ID == id {
			return &todos[i]
		}
	}
	return nil
}

// ToDoをIDで削除
func deleteTodoByID(id int) {
	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			return
		}
	}
}

// stringをintに変換する関数
func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
