package model

type VisitType string
type Location string

const (
	CLICK VisitType = "CLICK"
	DIRECT VisitType = "DIRECT"
)

const (
	GITHUB Location = "GITHUB"
	LINKEDIN Location = "LINKEDIN"
	HOME Location = "HOME"
	ABOUT Location = "ABOUT"
)

type Visit struct {
	From string `json:"from"`
	To string `json:"to"`
	Type VisitType `json:"type"`
	Location Location `json:"location"`
	Time int64 `json:"time"`
}