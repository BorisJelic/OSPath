package main

import (
	"fmt"
	"os"
)

func show(key string) (string, bool) {
	val, ok := os.LookupEnv(key)
	if !ok {
		fmt.Printf("%s not set\n", key)
	} else {
		fmt.Printf("%s=%s\n", key, val)
	}

	return val, ok
}

func main() {
	a, b := show("USER")

	fmt.Println(a)
	fmt.Println(b)
	//show("PWD")

}
