package handlers

type Employ struct {
	Emp_Id         string `json:"emp_id" yaml:"emp_id"`
	First_Name     string `json:"first_name" yaml:"first_name" `
	Second_Name    string `json:"second_name" yaml:"second_name"`
	Type           string `json:"types" yaml:"types"`
	Experience     string `json:"experience" yaml:"experience"`
	Default_Salary string `json:"default_salary" yaml:"default_salary"`
}
