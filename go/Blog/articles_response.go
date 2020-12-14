package Blog
type ArticlesResponse struct {
	PageCount int

	Articles []ArticleResponse `json:"Articles,omitempty"`
}