package main

import (
	"fmt"
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
	fmt.Println(rt)
	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		rv.Field(i).SetString(prms[f.Tag.Get("webcui")])
	}
	arg = rv.Interface()
	fmt.Println(arg)
	return arg
}

func main() {
	user := MapPosts(User{}).(User)
	fmt.Println(user)
}