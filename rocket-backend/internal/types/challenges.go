package types

type Challenge struct {
	ID     string `json:"id"`
	Text   string `json:"text"`
	Points int    `json:"points"`
}
