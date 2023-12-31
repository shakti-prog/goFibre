package functionsfolder

//Structs for different type of data coming from cassandra db

type userLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type serviceRequest struct {
	No          int64  `json:"no"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Status      string `json:"status"`
	Assignee    string `json:"assignee"`
	Reporter    string `json:"reporter"`
}

type TableData struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}
