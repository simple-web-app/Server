package Blog
type Comment struct {

	Date string `json:"date"`

	Content string `json:"content"`

	Author string `json:"author"`

	ArticleId int32 `json:"articleId"`
}