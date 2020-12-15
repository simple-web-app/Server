package swagger

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/boltdb/bolt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)
func CreateComments() {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//create
	err = db.Update(func(tx *bolt.Tx) error {
		a := tx.Bucket([]byte("Article"))
		c := a.Cursor()
		b := tx.Bucket([]byte("Comment"))

		if b == nil {
			b, err = tx.CreateBucket([]byte("Comment"))
			if err != nil {
				log.Fatal(err)
			}
		}
		if b != nil {
			var article Article
			for k, v := c.First(); k != nil; k, v = c.Next() {
				var comment Comment
				err := json.Unmarshal(v, &article)
				if err != nil {
					return err
				}
				fmt.Println(article.Id)
				var id int
				id = int(article.Id)
				filePath := "./data/" + strconv.Itoa(id) + "/comments"
				files, err := ioutil.ReadDir(filePath)
				if err != nil {
					log.Fatal(err)
				}

				for i := 1; i <= len(files); i++ {
					file, err := os.OpenFile(filePath+"/"+files[i-1].Name(), os.O_RDWR, 0666)
					buf := bufio.NewReader(file)
					user, err := buf.ReadString('\n')
					user = strings.TrimSpace(user)
					fmt.Println(user)

					time, err := buf.ReadString('\n')
					time = strings.TrimSpace(time)
					fmt.Println(time)

					var content string
					for {
						line, err := buf.ReadString('\n')
						line = strings.TrimSpace(line)
						content = content + line
						if err != nil {
							if err == io.EOF {
								fmt.Println("File read ok!")
								break
							} else {
								fmt.Println("Read file error!", err)
							}
						}
					}

					//timeStr := time.Now().Format("2006-01-02 15:04:05")
					comment = Comment{time, content, user, article.Id}
					fmt.Println(comment)
					vc, err := json.Marshal(comment)
					err = b.Put([]byte(strconv.Itoa(int(article.Id))+"_"+strconv.Itoa(i)), []byte(vc))
					if err != nil {
						log.Fatal(err)
					}
				}
			}
		} else {
			return errors.New("Table Comment doesn't exist")
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}
func CreateUser() {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("User"))
		if b == nil {
			//create table "xx" if not exits
			b, err = tx.CreateBucket([]byte("User"))
			if err != nil {
				log.Fatal(err)
			}
		}

		//insert rows
		for i := 0; i < 10; i++ {
			err := b.Put([]byte("user"+strconv.Itoa(i)), []byte("pass"+strconv.Itoa(i)))
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("user"+strconv.Itoa(i), "pass"+strconv.Itoa(i))
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}
func CreateTable() {
	db, err := bolt.Open("my.db", 0600, nil)
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

			filePath := "./data"
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