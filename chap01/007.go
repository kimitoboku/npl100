package main

import (
	"fmt"
)

func template(x, y, z interface{}) string {
	return fmt.Sprintf("%v時の%vは%v", x, y, z)
}

func main() {
	x := 12
	y := "気温"
	z := 22.4
	fmt.Printf("%+v\n", template(x, y, z))

}
