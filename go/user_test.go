package swagger

import (
	"fmt"
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




