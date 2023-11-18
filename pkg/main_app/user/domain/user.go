package domain //

type User struct {
	Firstname  string `json:"firstname"`
	Middlename string `json:"middlename"`
	Lastname   string `json:"lastname"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}
