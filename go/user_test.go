package swagger

import (
	"errors"
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"

	//my "github.com/simple-web-app/Server/go"
	"testing"
)

func TestCreateUser(t *testing.T) {
	fmt.Println()
	fmt.Println("Test for creating user...")

	test := []struct {
		name string
	}{
		{name: "testcase1: "},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(tt.name)
			CreateUser()
		})
	}
}

func signIn(username string, password string) Token {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("User"))
		if b != nil {
			v := b.Get([]byte(username))
			if ByteSliceEqual(v, []byte(password)) {
				return nil
			} else {
				return errors.New("Username and Password do not match")
			}
		} else {
			return errors.New("Username and Password do not match")
		}
	})

	if err != nil {
		log.Fatal(err)
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims

	if err != nil {
		log.Fatal(err)
	}

	tokenString, err := token.SignedString([]byte("my_secret_key"))
	if err != nil {
		log.Fatal(err)
	}

	response := Token{tokenString}
	return response
}

func TestSignIn(t *testing.T) {
	fmt.Println()
	fmt.Println("Test for User sign in...")

	test := []struct {
		username string
		password string
	}{
		{
			username: "user1",
			password: "pass1",
		},
		{
			username: "user5",
			password: "pass5",
		},
	}

	for _, tt := range test {
		t.Run(tt.username, func(t *testing.T) {
			fmt.Print("run testing of " +tt.username + "\n")
			signIn(tt.username, tt.password)
		})
	}
}




