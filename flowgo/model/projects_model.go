package model

type Contributor struct {
	Name string `json:"name"`
	Github string `json:"github"`
}

type Project struct {
	Name string `json:"name"`
	Repository string `json:"repository"`
	Description string `json:"description"`
	Links []string `json:"links"`
	Contributors []Contributor `json:"contributors"`
	Technologies []string `json:"technologies"`
	StartDate int64 `json:"startDate"`
	EndDate int64 `json:"endDate"`
}