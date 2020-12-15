package test

import (
	"fmt"
	my "github.com/simple-web-app/Server/go"
	"testing"
)

func TestCreateUser(t *testing.T) {
	fmt.Println("Test for creating user...")

	test := []struct {
		name string
	}{
		{name: "testcase1: "},
		{name: "testcase2: "},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(tt.name)
			my.CreateUser()
		})
	}
}


