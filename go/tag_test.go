package swagger

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"testing"
)

func getTagById(id int) {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()

	var tag Tag
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Tag"))
		fmt.Println(b)
		if b != nil {
			v := b.Get(itob(id))
			if v == nil {
				fmt.Println(id, " Tag Not Exists 1")
				return errors.New("Tag Not Exists 2 ")
			} else {
				_ = json.Unmarshal(v, &tag)
				return nil
			}
		} else {
			//fmt.Println("Tag Not Exists 3")
			return errors.New("Tag Not Exists 4")
		}
	})

	fmt.Println(tag.Name)
	fmt.Println("successful getTagById")
}


func getTags() {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()

	var tags Tags
	var tag Tag
	err = db.View(func(tx *bolt.Tx) error{
		b := tx.Bucket([]byte("Tag"))
		if b != nil{
			c := b.Cursor()
			for k, v := c.First(); k != nil; k, v = c.Next(){
				err = json.Unmarshal(v, &tag)
				//fmt.Println(err)
				//fmt.Println(v)
				if err != nil{
					return err
				}
				tags.Contents = append(tags.Contents, tag)
				//fmt.Println("")
			}
			return nil
		}else {
			return errors.New("tag Not Exists")
		}
	})

	for i:= 0; i < len(tags.Contents); i++ {
		fmt.Println(tags.Contents[i])
	}
	fmt.Println("successful getTags")
}

func addTag() {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()

	tagCount := 0
	fmt.Println("aaaaaaaaaaaaaa")
	db.View(func(tx *bolt.Tx) error{
		b := tx.Bucket([]byte("Tag"))
		b.ForEach(func(k, v []byte) error{
			tagCount = tagCount + 1
			return nil
		})
		return nil
	})
	fmt.Println(tagCount)

	var tag Tag
	tag.Name = "tag test"

	err = db.Update(func(tx *bolt.Tx) error{
		b, err := tx.CreateBucketIfNotExists([]byte("Tag"))
		if err != nil{
			return err
		}
		encoded, err := json.Marshal(tag)
		fmt.Println(tagCount)
		return b.Put(itob(tagCount+1), encoded)
	})

	if err != nil {
		log.Fatal(err)
	}
}

func Tags_t() {
	getTagById(1)
	getTags()
	//addTag()
}



func TestTags_t(t *testing.T) {
	fmt.Println()
	fmt.Println("Test for tag operations... ")

	test := []struct {
		name string
	}{
		{name: "testcase1: "},
		//{name: "testcase2: "},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(tt.name)
			Tags_t()
		})
	}
}