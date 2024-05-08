package models

type Employee struct {
	ID       int     `json:"id"`
	Name     string  `json:"name,omitempty"`
	Position string  `json:"position,omitempty"`
	Salary   float64 `json:"salary,omitempty"`
}

type GetEmployeeResponse struct {
	Employees     []Employee `json:"employees"`
	LastEvalKeyID int        `json:"last_eval_id,omitempty"`
}

type APIResponse struct {
	Message string `json:"message,omitempty"`
}

type EmployeeUpdateRequest struct {
	Position string  `json:"position,omitempty"`
	Salary   float64 `json:"salary,omitempty"`
}
