package webcui_test

import (
	"fmt"
	"github.com/melq/webcui-api"
	"log"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func TestMapPosts(t *testing.T) {
	type User struct {
		Name string `webcui:"name"`
		Age  string `webcui:"age"`
	}

	values := make(url.Values)
	values.Set("name", "Jiro")
	values.Set("age", "10")

	r, err := http.NewRequest(http.MethodPost, "https://example.com", strings.NewReader(values.Encode()))
	if err != nil {
		log.Fatalln("make request", err)
	}
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	user := &User{}
	err = webcui.MapPosts(user, r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user)
}

func ExampleMapPosts() {
	type User struct {
		Name string `webcui:"name"`
		Age  string `webcui:"age"`
	}

	user := &User{}
	var r *http.Request // put the appropriate instance in r

	_ = webcui.MapPosts(user, r) // returns err

	fmt.Println(user)
}

func TestExecCommand(t *testing.T) {
	cmd := "echo \"Taro Jiro\" \" Saburo \" \"Shiro\\\" \""
	res, err := webcui.ExecCommand(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(res))
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	bytes := []byte("a b c d \nURL: http://example.com")
	webcui.FmtAndWrite(bytes, w)
}

func TestFmtAndWrite(t *testing.T) {
	http.HandleFunc("/", handleRoot)

	fmt.Println("Listen..")
	log.Fatal("ListenAndServe", http.ListenAndServe(":8080", nil))
}
