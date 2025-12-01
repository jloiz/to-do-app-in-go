package types

type Task struct {
	TaskId   int     `json:"taskId"`
	TaskBody string  `json:"taskBody"`
	Status   string  `json:"status"`
}

type ErrorRes struct {
	Error string `json:"error"`
}