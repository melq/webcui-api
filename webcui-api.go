package webcui_api

import (
	"reflect"
)

type User struct {
	Name 	string	`webcui:"name"`
	Age 	string	`webcui:"age"`
}

func MapPosts(arg interface{}) interface{} {
	prms := map[string]string{"name": "Jiro", "age": "10"}

	rv := reflect.New(reflect.TypeOf(arg)).Elem()
	rt := reflect.TypeOf(arg)
	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		rv.Field(i).SetString(prms[f.Tag.Get("webcui")])
	}
	arg = rv.Interface()
	return arg
}
