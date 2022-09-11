package api

type BlogPost struct {
	Title string `json:"title"`
	Body  string `json:"md_body"`
}
