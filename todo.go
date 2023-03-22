package todo

type ToDo struct {
	Name string `json:"name"`
	Done bool   `json:"done,string"`
	// Deadline time.Time
}

type List struct {
	todos []ToDo
}

func (t *ToDo) AddTask() {

}
func (t *ToDo) DeleteTask() {

}
func (t *ToDo) DoneTask() {

}
func (t *ToDo) UpdateTask() {

}
