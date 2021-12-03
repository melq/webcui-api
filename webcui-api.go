package webcui

import (
	"net/http"
	"reflect"
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
