package yousign

import (
	"fmt"
	"testing"
)

func TestUserAll(t *testing.T) {
	users, _, err := client.User.All()

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("user found:", len(users))
}
