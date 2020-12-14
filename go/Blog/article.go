package Blog
type Article struct {

	Id int32 `json:"id"`

	Name string `json:"name"`

	Tags []Tag `json:"tags,omitempty"`

	Date string `json:"date,omitempty"`

	Content string `json:"content"`
}