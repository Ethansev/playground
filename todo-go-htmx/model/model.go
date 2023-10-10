package model

type Todo struct {
	Id   uint64 `json:"id"`
	Todo string `json:"todo"`
	Done bool   `json:"done"`
}

func CreateTodo(todo string) error {
	statement := "INSERT INTO todos (todo, done) VALUES ($1, $2);"

	_, err := db.Query(statement, todo, false)

	return err
}

func GetAllTodos() ([]Todo, error) {
	todos := []Todo{}

	statement := "SELECT id , todo, done FROM todos;"

	rows, err := db.Query(statement)
	if err != nil {
		return todos, err
	}

	defer rows.Close()

	for rows.Next() {
		var title string
		var done bool
		var id uint64

		err := rows.Scan(&id, &title, &done)
		if err != nil {
			return todos, err
		}
		todo := Todo{
			Id:   id,
			Todo: title,
			Done: done,
		}

		todos = append(todos, todo)
	}

	return todos, err
}

func GetTodo(id uint64) (Todo, error) {
	statement := "SELECT id, todo, done FROM todos WHERE id = $1;"

	todo := Todo{}
	todo.Id = id

	row, err := db.Query(statement, id)
	if err != nil {
		return todo, err
	}

	for row.Next() {
		var title string
		var done bool

		err := row.Scan(&id, &title, &done)
		if err != nil {
			return todo, err
		}

		todo.Todo = title
		todo.Done = done
	}

	return todo, err
}

func MarkDone(id uint64) error {
	todo, err := GetTodo(id)
	if err != nil {
		return err
	}

	statement := "UPDATE todos SET done = $2 WHERE id = $1;"
	_, err = db.Query(statement, id, !todo.Done)

	// if todo.Done {
	// 	_, err = db.Query(statement, id, false)
	// } else {
	// 	_, err = db.Query(statement, id, true)
	// }

	return err
}

func Delete(id uint64) error {
	statement := "DELETE FROM todos WHERE id = $1;"

	_, err := db.Exec(statement, id)
	return err
}
