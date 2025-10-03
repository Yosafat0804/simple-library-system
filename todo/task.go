package todo

// Struct untuk menyimpan 1 Task
type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}