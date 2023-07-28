package hireddto

import "juna/models"

type CreateHiredResponse struct {
	Title        string      `json:"title"`
	Description  string      `json:"description"`
	StartProject string      `json:"startProject"`
	EndProject   string      `json:"endProject"`
	Price        int         `json:"price"`
	OrderBy      models.User `json:"orderBy"`
	OrderTo      models.User `json:"orderTo"`
}

type HiredResponse struct {
	Hired interface{} `json:"hired"`
}
