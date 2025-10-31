package dto

type ContextRequest struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type ShowAllContexts struct {
	Limit int `json:"limit"`
}

type UpdateContextRequest struct {
	Name *string `json:"name"`
	Desc *string `json:"desc"`
	IsHidden *bool `json:"ishidden"`
}