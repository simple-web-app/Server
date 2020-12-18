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
	"time"
)
func CreateComments() {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//create
	err = db.Update(func(tx *bolt.Tx) error {
		//open the article database and comment database
		a := tx.Bucket([]byte("Article"))
		c := a.Cursor()
		b := tx.Bucket([]byte("Comment"))
		//if comment database isn't exists
		if b == nil {
			b, err = tx.CreateBucket([]byte("Comment"))
			if err != nil {
				log.Fatal(err)
			}
		} else{
			return nil
		}
		//follow article to read comments
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
				if err == nil {
						b1 := tx.Bucket([]byte("Article"))
						if b1 != nil{
							c := b1.Cursor()
							c.Seek(itob(id))
							errs := c.Delete()
							article.CommentsNum = article.CommentsNum + int32(len(files))
							encode, errs := json.Marshal(article)
							fmt.Println(errs)
							b1.Put(itob(id), encode)
							if errs != nil{
								return errors.New("Delete article failed")
							}
						} else{
							return errors.New("Ariticle Not Exists")
						}
					for i := 1; i <= len(files); i++ {
						file, errs := os.OpenFile(filePath+"/"+files[i-1].Name(), os.O_RDWR, 0666)
						fmt.Println(errs)
						buf := bufio.NewReader(file)
						user, errs := buf.ReadString('\n')
						user = strings.TrimSpace(user)
						fmt.Println(user)
						time, errs := buf.ReadString('\n')
						time = strings.TrimSpace(time)
						fmt.Println(time)
						var content string
						for {
							line, errs := buf.ReadString('\n')
							line = strings.TrimSpace(line)
							content = content + line
							if errs != nil {
								if errs == io.EOF {
									fmt.Println("File read ok!")
									break
								} else {
									fmt.Println("Read file error!", errs)
								}
							}
						}
						//timeStr := time.Now().Format("2006-01-02 15:04:05")
						comment = Comment{time, content, user, article.Id}
						fmt.Println(comment)
						vc, errs := json.Marshal(comment)
						errs = b.Put([]byte(strconv.Itoa(int(article.Id))+"_"+strconv.Itoa(i)), []byte(vc))
						if errs != nil {
							log.Fatal(errs)
						}
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
	// open database
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// if not exists User bucket
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("User"))
		if b == nil {
			b, err = tx.CreateBucket([]byte("User"))
			if err != nil {
				log.Fatal(err)
			}
		} else{
			return nil
		}
	// add origin user1~user10
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
		// open database
		b := tx.Bucket([]byte("Article"))
		if b == nil {
			//create table "xx" if not exits
			b, err = tx.CreateBucket([]byte("Article"))
			if err != nil {
				log.Fatal(err)
			}
		} else{
			return nil
		}
		// create origin database and load local data
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
					fmt.Println("failed", err)
					return err
				}

				title := articleName[:len(articleName)-3]
				var timeUnix int64 = int64(i) * 100 + 1608120901
				article = Article{int32(i), title, tags, (time.Unix(timeUnix, 0)).Format("2006-01-02 15:04:05"), string(content), 0}
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


func CreateTag() {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		// if not exists Tag bucket
		b := tx.Bucket([]byte("Tag"))
		if b == nil {
			b, err = tx.CreateBucket([]byte("Tag"))
			if err != nil {
				log.Fatal(err)
			}
		} else{
			return nil
		}
		// add origin tag1 and tag2
		var tag Tag
		tag.Name = "tag1"
		v, err := json.Marshal(tag)
		err = b.Put([]byte(itob(1)), v)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(itob(1), "tag"+strconv.Itoa(1))
		/*
		tag.Name = "tag2"
		v, err = json.Marshal(tag)
		err = b.Put([]byte(strconv.Itoa(1)), v)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(strconv.Itoa(2), "tag"+strconv.Itoa(2))
		*/

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}

