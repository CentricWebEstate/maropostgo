package marogo

const API_URL string = "https://api.maropost.com/accounts/"

type Maropost struct {
	AuthToken string
	Account   string
}

type Contact struct {
	Maropost
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Fax         string `json:"fax"`
	CustomField map[string]interface{}
	Subscribe   bool `json:"subscribe"`
}
