package representations

type Page struct {
	Title string `json:"title"`
	Story []string `json:"story"`
	Options []Option `json:"options"`
}