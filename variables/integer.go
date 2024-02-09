package variables

import (
	"fmt"
)

func PrintIntegers() {
	var x int        // variable de declaracion
	y := int32(1243) // variable de asignacion
	z := int64(652345)
	fmt.Println("X:", x)
	fmt.Println("Y:", y)
	fmt.Println("Z:", z)
}
