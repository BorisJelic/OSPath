package main

import( 
	"fmt"

)

func change(v *int) {
	*v = 50
}

func modify(sls []int) {
	sls[0] = 1029094
}

func main() {
	b := 10000

	var a = &b 
	
	fmt.Printf("Type of a is %T\n", a)
	fmt.Println("value of b is", *a)
	*a++
	fmt.Println("value of b is", *a)

   change(a)
   fmt.Println("value of b is", a)

	abc := [...]int{100, 2304, 234, 50, 90,293,39485}

	modify(abc[:])
	fmt.Println(abc)

}