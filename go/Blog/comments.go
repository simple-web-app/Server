package Blog
type Comments struct {
	PageCount int

	Contents []Comment `json:"contents,omitempty"`
}