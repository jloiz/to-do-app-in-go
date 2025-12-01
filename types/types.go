package types

type Task struct {
	TaskId   int
	TaskBody string
	Status   string
}

type ErrorRes struct {
	Error string
}