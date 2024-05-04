package domain

type Todo struct {
	ID          string `json:"id" example:"1"`
	Name        string `json:"name" example:"Example name"`
	Description string `json:"description" example:"Example description"`
	Done        bool   `json:"done"`
}

func (Todo) TableName() string {
	return "todo"
}
