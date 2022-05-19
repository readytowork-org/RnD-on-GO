//packages
package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var a int = 5
	b := "ram"

	fmt.Println("Hello World", a, b)
	if a > 5 {
		fmt.Println("a is greater")
	} else if a < 5 {
		fmt.Println("5 is greater")
	} else {
		fmt.Println("They are equal")
	}

	// Arrays
	c := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(c)
	slicedArray := c[2:8]
	appendedArray := append(c, 10)
	fmt.Println(slicedArray, appendedArray)

	//sorting an array
	newArray := []int{345, 78, 123, 10, 76, 2, 567, 5}
	sort.Ints(newArray)
	fmt.Println("Sorted array: ", newArray)

	//finding index in a string
	fmt.Println("Index of hello is: ", strings.Index("This is a boy. Say hello to the boy.", "hello"))

	// Map
	d := make(map[string]string)
	d["first"] = "first value"
	d["second"] = "second value"
	d["third"] = "third value"
	d["fourth"] = "fourth value"
	d["fifth"] = "fifth value"
	d["sixth"] = "sixth value"

	// delete from Map
	delete(d, "sixth")
	fmt.Println(d)

	// for Loop
	for i := 0; i < 5; i++ {
		fmt.Println("The values are", i)
	}
	for key, value := range d {
		fmt.Println("key: ", key, "value: ", value)
	}
	addition(5, 2)
}
func addition(a int, b int) {
	sum := a + b
	fmt.Println(sum)
}
