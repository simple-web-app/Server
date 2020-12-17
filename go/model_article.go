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

type Article struct {

	Id int32 `json:"id"`

	Name string `json:"name"`

	Tags []Tag `json:"tags,omitempty"`

	Date string `json:"date,omitempty"`

	Content string `json:"content"`

	//CommentsNum int32 `json:"commentNum, omitempty"`
}
