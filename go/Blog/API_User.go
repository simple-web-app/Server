package Blog

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"
	"github.com/boltdb/bolt"
	"github.com/dgrijalva/jwt-go"
	"log"
)

func SignIn(w http.ResponseWriter, r *http.Request){
	db, err := bolt.Open("my.db", 0600, nil);
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()
	u, err :=url.Parse(r.URL.String())
	if err != nil{
		log.Fatal(err)
	}
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Printf(m)
	var user User
	user.Username = m["username"][0]
	user.Password = m["password"][0]
	fmt.Println(user.Username, user.Password)
	err = db.View(func(tx *bolt.Tx) error{
		b := tx.Bucket([]byte("User"))
		if b != nil{
			v := b.Get([]byte(user.Username))
			if ByteSliceEqual(v, []byte(user.Password)){
				return nil
			}
		}
	})

}

func ByteSliceEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	if (a == nil) != (b == nil) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
