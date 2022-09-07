package note

import (
	"fmt"
	"unsafe"
)

type s1 struct{ Val int }
type s2 struct{ Val int }

func note() {
	var a interface{}
	v := a.(int)
	fmt.Println(v)

	var p1 *s1
	var p2 *s2 = (*s2)(p1)
	fmt.Println(p2)

	var d *float32
	//compile fail
	//var x *int = (*int)(d)
	var x *int = (*int)(unsafe.Pointer(d))
	fmt.Println(x)
}
