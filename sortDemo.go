package sortdemo

import(
	"fmt"
	"sort"
)

func sortdemo() {

	people := []struct {
			Name string
			Age int
	}{
		{"Boris", 15},
		{"David", 10},
		{"John", 55},
	}


sort.Slice(people, func(i, j int) bool { return people[i].Name < people[j].Name})
fmt.Println("By name: ", people)

sort.Slice(people, func(i, j int) bool { return people[i].Age < people[j].Age})
fmt.Println("By age:", people)


}