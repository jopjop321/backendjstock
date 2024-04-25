package forms

type Register struct {
	Name	string `json:"name"`
	Last_name string `json:"last_name"`
	Nickname string `json:"nickname"`
	ID string `json:"id"`
	Password string `json:"password"`
}

type Login struct {
	ID string `json:"id"`
	Password string `json:"password"`
}

type Employee_information struct {
	Employee_ID int `json:"employee_id"`
	Name	string `json:"name"`
	Last_name string `json:"last_name"`
	Nickname string `json:"nickname"`	
}

type Check_information struct {
	Name	string `json:"name"`
	Last_name string `json:"last_name"`
	Nickname string `json:"nickname"`	
}

type Employee_profile struct {
	Employee_ID int `json:"employee_id"`
	Name	string `json:"name"`
	Last_name string `json:"last_name"`
	Nickname string `json:"nickname"`
	Position string `json:"position"`	
}

type Update_Status struct {
	Employee_ID int `json:"employeeid"`
	Position string `json:"position"`
}

type Profile struct {
	Employee_ID int `json:"employeeid"`
	Nickname string `json:"nickname"`
	Position string `json:"position"`
}

