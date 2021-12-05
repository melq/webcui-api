package webcui

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"regexp"
	"strings"
)

// MapPosts リクエストのPOSTパラメータを引数の構造体ポインタの実体にマッピングします
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
