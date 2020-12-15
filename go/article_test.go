package swagger

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/boltdb/bolt"
	"io/ioutil"
	"strconv"
	"testing"

	//my "github.com/simple-web-app/Server/go"
	"log"
)

//getArticleById use for testing the func in api_article
func getArticleById(id int) {
	db, err := bolt.Open("../my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var article Article
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Article"))
		fmt.Println(b)
		if b != nil {
			v := b.Get(itob(id))
			if v == nil {
				fmt.Println(id, " Article Not Exists")
				return errors.New("Article Not Exists")
			} else {
				_ = json.Unmarshal(v, &article)
				return nil
			}
		} else {
			fmt.Println("Article Not Exists")
			return errors.New("Article Not Exists")
		}
	})

}

//getArticles used for test the func in api_article
func getArticles(p int) {
	db, err := bolt.Open("../.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//display 10 articles per page
	IdIndex := (p-1)*10 + 1
	var articles ArticlesResponse
	var article ArticleResponse
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Article"))
		if b != nil {
			c := b.Cursor()
			k, v := c.Seek(itob(IdIndex))
			if k == nil {
				fmt.Println("Page is out of index")
				return errors.New("Page is out of index")
			}
			key := binary.BigEndian.Uint64(k)
			fmt.Print(key)
			if int(key) != IdIndex {
				fmt.Println("Page is out of index")
				return errors.New("Page is out of index")
			}
			count := 0
			var ori_artc Article
			for ; k != nil && count < 10; k, v = c.Next() {
				err = json.Unmarshal(v, &ori_artc)
				if err != nil {
					return err
				}
				article.Id = ori_artc.Id
				article.Name = ori_artc.Name
				articles.Articles = append(articles.Articles, article)
				count = count + 1
			}
			return nil
		} else {
			return errors.New("Article Not Exists")
		}
	})
	for i := 0; i < len(articles.Articles); i++ {
		fmt.Println(articles.Articles[i])
	}
}

//deleteArticleById used for testing the func in api_article
func deleteArticleById(id int) {
	//connect to database
	db, err := bolt.Open("../my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//delete the article by ID
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Article"))
		if b != nil {
			c := b.Cursor()
			c.Seek(itob(id))
			err := c.Delete()
			if err != nil {
				//fmt.Println("Delete article failed")
				//log.Fatal(err)
				return errors.New("Delete article failed")
			}
		} else {
			//fmt.Println("Article Not Exists")
			return errors.New("Article Not Exists")
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Successfully Delete article ", id)
}

func createTable() {
	db, err := bolt.Open("../my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Article"))
		if b == nil {
			//create table "xx" if not exits
			b, err = tx.CreateBucket([]byte("Article"))
			if err != nil {
				log.Fatal(err)
			}
		}
		if b != nil {
			var article Article
			var tags []Tag
			tags = append(tags, Tag{"tag1"})
			tags = append(tags, Tag{"tag2"})

			filePath := "../data"
			files, err := ioutil.ReadDir(filePath)
			if err != nil {
				log.Fatal(err)
			}
			for i := 1; i <= len(files); i++ {
				path := filePath + "/" + strconv.Itoa(i)
				fileInfoList, err := ioutil.ReadDir(path)
				var articleName string
				for i := 0; i < len(fileInfoList); i++ {
					if fileInfoList[i].IsDir() == false {
						articleName = fileInfoList[i].Name()
					}
				}
				if err != nil {
					log.Fatal(err)
				}
				content, err := ioutil.ReadFile(path + "/" + articleName)
				if err != nil {
					fmt.Println("获取失败", err)
					return err
				}

				//fmt.Println("文本内容为:", string(content))

				title := articleName[:len(articleName)-3]
				article = Article{int32(i), title, tags, "2020", string(content)}
				v, err := json.Marshal(article)
				//insert rows
				err = b.Put(itob(i), v)
				if err != nil {
					log.Fatal(err)
				}
			}
		} else {
			return errors.New("Table Article doesn't exist")
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}

func Articles() {
	createTable()
	getArticleById(1)
	getArticles(1)

	deleteArticleById(1)
	//getArticleById(1)

}

func TestArticles(t *testing.T) {
	fmt.Println()
	fmt.Println("Test for article operations... ")

	test := []struct {
		name string
	}{
		{name: "testcase1: "},
		//{name: "testcase2: "},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(tt.name)
			Articles()
		})
	}
}
