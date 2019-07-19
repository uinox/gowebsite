package main

import (
	"io"
	"os"
	"fmt"
	"reflect"
)

func main()  {
	var w io.Writer = os.Stdout
	fmt.Println(reflect.TypeOf(w))
	fmt.Printf("%T\n", 3)

	fmt.Println("----")
	v := reflect.ValueOf(3)
	fmt.Println(v)
	fmt.Printf("%v\n", v)
	fmt.Println(v.String())

	t := v.Type()
	fmt.Println(t.String())

	x := reflect.ValueOf((3))
	y := x.Interface()
	i := y.(int)
	fmt.Println(y, i)
}
