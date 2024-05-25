package domain

type Todo struct {
	ID          string `json:"id" example:"4b61e0a8-1fe0-4c7f-97bf-f3e9c4e86c3a"`
	Name        string `json:"name" example:"Example name"`
	Description string `json:"description" example:"Example description"`
	Done        bool   `json:"done"`
}

func (Todo) TableName() string {
	return "todo"
}

type PaginatedTodo struct {
	Data []Todo `json:"data"`
	Page int    `json:"page"`
}

func NewPaginatedTodo(data []Todo, page int) *PaginatedTodo {
	return &PaginatedTodo{
		Data: data,
		Page: page,
	}
}
