package webcui

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"reflect"
	"regexp"
	"strings"
)

// MapPosts maps each POST value in r to a field in struct *arg that has a value of the webcui tag with the same name as the POST key.
//
// Example
// type User struct {
//	Name string `webcui:"name"`
//	Age  string `webcui:"age"`
// }
//
// user := &User{}
// _ = MapPosts(user, r) // returns err
// fmt.Println(user)
func MapPosts(arg interface{}, r *http.Request) error {
	if reflect.TypeOf(arg).Kind() != reflect.Ptr || reflect.ValueOf(arg).Elem().Kind() != reflect.Struct {
		return errors.New("arg is not Ptr to Struct")
	}

	vp := reflect.ValueOf(arg)
	rt := reflect.Indirect(vp).Type()

	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		vp.Elem().Field(i).SetString(r.FormValue(f.Tag.Get("webcui")))
	}

	return nil
}

// ExecCommand returns result of calling the CUI based on cmd
func ExecCommand(cmd string) ([]byte, error) {
	s := strings.Split(cmd, " ")
	name := s[0]
	args := s[1:]
	res, err := exec.Command(name, args...).Output()
	if err != nil {
		return []byte(""), err
	}

	return res, nil
}

// FmtAndWrite formats bytes to a format suitable for webcui and writes them to the body of w
func FmtAndWrite(bytes []byte, w http.ResponseWriter) {
	str := string(bytes)
	strSlice := strings.Split(str, "\n")

	reg := regexp.MustCompile(`(https?://[\w/:%#$&?()~.=+\-]+$)`)
	for _, str := range strSlice {
		if reg.Match([]byte(str)) {
			str = reg.ReplaceAllString(str, "<a href=\"$1\">$1</a>")
		}
		str += "<br>"
		_, err := fmt.Fprintf(w, str)
		if err != nil {
			log.Println(err)
		}
	}
}
