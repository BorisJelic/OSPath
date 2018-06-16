package main

import (
	"sort"
	"fmt"
	"os"
	"io/ioutil"
	"log"
)

func show(key string) (string, bool) {
	val, ok := os.LookupEnv(key)

	if !ok {
		fmt.Printf("%s not set\n", key)
	} else {
		fmt.Printf("%s=%s\n", key, val)
	}

	fmt.Printf("GET PID %v \n\n", os.Getuid())

	return val, ok
}


type fileNames struct {
	Filename string
	Filesize int64
}



func main() {
//	a, b := show("USER")

	//fmt.Println(a)
	//fmt.Println(b)
	//show("PWD")

files, err := ioutil.ReadDir("/Users/boris/Downloads")
if err != nil {
	log.Fatal(err)
}

ctx := context.Background()


var m map[string]int64

m = make(map[string]int64)
var keys []fileNames


for _, file := range files {
	
	name := file.Name()
	size := file.Size()/1024/1024
	m[name] = size
}

for k, v := range m {
	fn := fileNames{k,v}
	keys = append(keys, fn)
}
sort.Slice(keys, func(i, j int) bool { return keys[i].Filesize < keys[j].Filesize } )


for _, v := range keys {
	fmt.Printf("File name: %v and filesize: %v MB \n", v.Filename, v.Filesize)
}



}
