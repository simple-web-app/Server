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

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"github.com/boltdb/bolt"
)

func DeleteArticleById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("call deleter")
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()
	articleId := strings.Split(r.URL.Path, "/")[4]
	Id, err := strconv.Atoi(articleId)
	fmt.Println(Id)
	if err != nil{
		response := ErrorResponse{"Wrong ArticledId"}
		JsonResponse(response, w, http.StatusBadRequest)
		return
	}
	err = db.Update(func(tx *bolt.Tx)error{
		b := tx.Bucket([]byte("Article"))
		if b != nil{
			c := b.Cursor();
			c.Seek(itob(Id))
			err := c.Delete()
			if err != nil{
				return errors.New("Delete article failed")
			}
		} else{
			return errors.New("Ariticle Not Exists")
		}
		return nil
	})
	if err != nil{
		response := ErrorResponse{err.Error()}
		JsonResponse(response, w, http.StatusNotFound)
		return
	}
	JsonResponse("Success",  w, http.StatusOK)
}

func GetArticleById(w http.ResponseWriter, r *http.Request) {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()
	articleId := strings.Split(r.URL.Path, "/")[4]
	Id, err := strconv.Atoi(articleId)
	fmt.Println(Id)
	if err != nil{
		reponse := ErrorResponse{"Wrong ArticleId"}
		JsonResponse(reponse, w, http.StatusBadRequest)
		return
	}
	var article Article
	err = db.View(func(tx *bolt.Tx) error{
		b := tx.Bucket([]byte("Article"))
		if b != nil{
			v := b.Get(itob(Id))
			if v == nil{
				return errors.New("Article Not Exists")
			}else{
				_ = json.Unmarshal(v, &article)
				return nil
			}
		} else{
			return errors.New("Article Not Exists")
		}
	})
	if err != nil{
		response := ErrorResponse{err.Error()}
		JsonResponse(response, w, http.StatusNotFound)
		return
	}
	JsonResponse(article, w, http.StatusOK)
}

func GetArticles(w http.ResponseWriter, r *http.Request) {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()
	u, err := url.Parse(r.URL.String())
	if err != nil{
		log.Fatal(err)
	}
	m, _ := url.ParseQuery(u.RawQuery)
	page := m["page"][0]
	IdIndex, err := strconv.Atoi(page)
	pageCount := 0
	db.View(func(tx *bolt.Tx) error{
		b := tx.Bucket([]byte("Article"))
		b.ForEach(func(k, v []byte) error{
			pageCount = pageCount + 1
			return nil
		})
		return nil
	})
	fmt.Println(pageCount)
	IdIndex = (IdIndex - 1) * 10 + 1
	var articles ArticlesResponse
	var article ArticleResponse
	err = db.View(func(tx *bolt.Tx) error{
		articles.PageCount = pageCount
		b := tx.Bucket([]byte("Article"))
		var k, v []byte
		if b != nil{
			c := b.Cursor()
			k, v = c.First()
			err = json.Unmarshal(v, &article)
			for i := 1; i < IdIndex; i++{
				k, v = c.Next()
				err = json.Unmarshal(v, &article)
				fmt.Println(article.Id)
				if k == nil{
					return errors.New("Page is out of index")
				}
			}
			count := 0
			for ; k != nil && count < 10; k, v = c.Next(){
				err = json.Unmarshal(v, &article)
				if err != nil{
					return err
				}
				articles.Articles = append(articles.Articles, article)
				count = count + 1
			}
			return nil
		} else{
			return errors.New("Article Not Exists")
		}
	})
	if err != nil{
		response := ErrorResponse{err.Error()}
		JsonResponse(response, w, http.StatusNotFound)
	}
	json, err := json.Marshal(articles)
	fmt.Println(string(json))
	JsonResponse(articles, w, http.StatusOK)
}
