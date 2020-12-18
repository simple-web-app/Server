/*
 * simple web blog
 *
 * Redefine the Tag as resource.
 *
 * API version: 1.0.0
 * Contact: lurui7@mail2.sysu.edu.cn
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"net/http"
	"strconv"
	"strings"
)



func GetTagById(w http.ResponseWriter, r *http.Request) {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()

	tagId := strings.Split(r.URL.Path, "/")[4]
	Id, err := strconv.Atoi(tagId)
	if err != nil{
		response := ErrorResponse{"Wrong TagId"}
		JsonResponse(response, w, http.StatusBadRequest)
		return
	}

	var tag Tag
	err = db.View(func(tx *bolt.Tx) error{
		b := tx.Bucket([]byte("Tag"))
		if b != nil{
			fmt.Println(Id)
			v := b.Get(itob(Id))
			if v == nil{
				return errors.New("tag Not Exists 1")
			}else{
				_ = json.Unmarshal(v, &tag)
				return nil
			}
		} else{
			return errors.New("tag Not Exists 2")
		}
	})
	if err != nil{
		response := ErrorResponse{err.Error()}
		JsonResponse(response, w, http.StatusNotFound)
		return
	}
	JsonResponse(tag, w, http.StatusOK)
}

func GetTags(w http.ResponseWriter, r *http.Request) {
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
			}
			return nil
		}else {
			return errors.New("tag Not Exists")
		}
	})

	if err != nil{
		response := Response404{err.Error()}
		JsonResponse(response, w, http.StatusNotFound)
		return
	}
	JsonResponse(tags, w, http.StatusOK)
}


func AddTag(w http.ResponseWriter, r *http.Request) {
	fmt.Println("trying to add tag")
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()

	tagCount := 0
	db.View(func(tx *bolt.Tx) error{
		b := tx.Bucket([]byte("Tag"))
		b.ForEach(func(k, v []byte) error{
			tagCount = tagCount + 1
			return nil
		})
		return nil
	})
	fmt.Println(tagCount)

	tag := Tag{
		Name: "",
	}

	err = json.NewDecoder(r.Body).Decode(&tag)
	if err != nil || (tag.Name == "") {
		w.WriteHeader(http.StatusBadRequest)
		if err != nil{
			response := Response404{err.Error()}
			JsonResponse(response, w, http.StatusBadRequest)
		} else{
			response := ErrorResponse{"There is no name"}
			JsonResponse(response, w, http.StatusBadRequest)
		}
		return
	} else{
		fmt.Println("tag added:", tag)
	}

	err = db.Update(func(tx *bolt.Tx) error{
		b, err := tx.CreateBucketIfNotExists([]byte("Tag"))
		if err != nil{
			return err
		}
		encoded, err := json.Marshal(tag)
		fmt.Println(tagCount)
		return b.Put(itob(tagCount+1), encoded)
	})
	if err != nil{
		response := Response404{err.Error()}
		JsonResponse(response, w, http.StatusBadRequest)
		return
	}
	JsonResponse(tag, w, http.StatusOK)
}



