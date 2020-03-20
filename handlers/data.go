package handlers

type Employ struct {
	E_Id       string `json:"emp_id" yaml:"emp_id"`
	F_Name     string `json:"first_name" yaml:"first_name" `
	S_Name     string `json:"second_name" yaml:"second_name"`
	Type       string `json:"types" yaml:"types"`
	Exper      string `json:"experience" yaml:"experience"`
	Def_Salary string `json:"default_salary" yaml:"default_salary"`
}
