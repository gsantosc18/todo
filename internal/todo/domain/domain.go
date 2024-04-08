package domain

type Todo struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type Response struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}
