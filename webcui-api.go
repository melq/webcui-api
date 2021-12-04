package webcui

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"regexp"
	"strings"
)

func MapPosts(arg interface{}, r *http.Request) interface{} {
	rv := reflect.New(reflect.TypeOf(arg)).Elem()
	rt := reflect.TypeOf(arg)
	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		rv.Field(i).SetString(r.FormValue(f.Tag.Get("webcui")))
	}
	arg = rv.Interface()
	return arg
}

// FmtAndWrite Responseのbodyを整形、Responseに書き込みする関数
func FmtAndWrite(bytes []byte, w http.ResponseWriter) {
	str := string(bytes)
	strSlice := strings.Split(str, "\n")

	reg := regexp.MustCompile(`https?://[\w/:%#$&?()~.=+\-]+$`)
	for _, str := range strSlice {
		if reg.Match([]byte(str)) {
			str = fmt.Sprintf("<a href=\"%s\">%s</a>", str, str)
		}
		str += "<br>"
		_, err := fmt.Fprintf(w, str)
		if err != nil {
			log.Println(err)
		}
	}
}

// 必須タグ等で入力コマンドを自動で作成する関数
