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
func MapPosts(dest interface{}, r *http.Request) error {
	if reflect.TypeOf(dest).Kind() != reflect.Ptr || reflect.ValueOf(dest).Elem().Kind() != reflect.Struct {
		return errors.New("dest is not Ptr to Struct")
	}

	vp := reflect.ValueOf(dest)
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
	s = s[1:]

	args := make([]string, 0)
	isInDoubleQuote := false
	for _, v := range s {
		if !isInDoubleQuote || len(args) == 0 {
			args = append(args, v)
		} else {
			args[len(args)-1] += " " + v
		}

		if v == "\"" {
			isInDoubleQuote = !isInDoubleQuote
		} else {
			if !isInDoubleQuote && strings.HasPrefix(v, "\"") {
				isInDoubleQuote = true
			}
			if isInDoubleQuote && strings.HasSuffix(v, "\"") && !strings.HasSuffix(v, "\\\"") {
				isInDoubleQuote = false
			}
		}
	}
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
