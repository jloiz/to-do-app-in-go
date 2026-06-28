package types

type Task struct {
	TaskId   string     `json:"taskId"`
	TaskBody string     `json:"taskBody"`
	Status   string     `json:"status"`
}

type TaskRequest struct {
	TaskBody string  `json:"taskBody"`
	Status   string  `json:"status"`
}

type SuccessResponse struct {
	TaskId string `json:"taskId"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type WriteTaskResponse struct {
	// ToDo: change to UUid
	TaskId string `json:"taskId,omitempty"`
	Error string `json:"error,omitempty"`
}

