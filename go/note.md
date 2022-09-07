```go
package note
import (
	"fmt"
	"unsafe"
)
func note() {
	var a interface{}
	v,ok := a.(int)
	if ok {
		fmt.Println(v)
    }

	var d *float32
	//compile fail
	//var x *int = (*int)(d) 
	var x *int = (*int)(unsafe.Pointer(d))
	fmt.Println(x)
}
```
