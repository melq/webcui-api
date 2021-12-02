package webcui_api

import (
	"fmt"
	"testing"
)

func TestMapPosts(t *testing.T) {
	user := MapPosts(User{}).(User)
	fmt.Println(user)
}
