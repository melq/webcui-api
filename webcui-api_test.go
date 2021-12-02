package webcui_api

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func TestMapPosts(t *testing.T) {
	type User struct {
		Name 	string	`webcui:"name"`
		Age 	string	`webcui:"age"`
	}

	values := make(url.Values)
	values.Set("name", "Jiro")
	values.Set("age", "10")

	r, err := http.NewRequest(http.MethodPost, "https://example.com", strings.NewReader(values.Encode()))
	if err != nil {
		log.Fatalln("make request", err)
	}
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	user := MapPosts(User{}, r).(User)
	fmt.Println(user)
}
