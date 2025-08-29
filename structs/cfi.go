package structs

type Cfi struct {
	Id              int    `json:"id"`
	Name            string `json:"name"`
	Website         string `json:"website"`
	Email           string `json:"email"`
	Stage           string `json:"stage"`
	OrgName         string `json:"org"`
	VisionStatement string `json:"vision_statement"`
}
