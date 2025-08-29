package structs

type Ramp struct {
	Id             int             `json:"id"`
	Name           string          `json:"name"`
	Url            string          `json:"url"`
	Email          string          `json:"email"`
	Regions        []Region        `json:"regions"`
	Assets         []Asset         `json:"assets"`
	PaymentMethods []PaymentMethod `json:"payment_methods"`
	Approval       bool            `json:"approval"`
}

type Region struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Asset struct {
	Id      int    `json:"id"`
	Ticker  string `json:"ticker"`
	Address string `json:"address"`
}

type PaymentMethod struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
