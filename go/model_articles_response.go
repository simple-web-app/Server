/*
 * simple web blog
 *
 * Simple Web Blog
 *
 * API version: 1.0.0
 * Contact: nanzh@mail2.sysu.edu.cn
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ArticlesResponse struct {
	PageCount int
	Articles []ArticleResponse `json:"Articles,omitempty"`
}
