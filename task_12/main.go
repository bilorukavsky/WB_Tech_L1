/*
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
*/
package main

import "fmt"

func createSet(strings []string) map[string]bool {
	set := make(map[string]bool)
	for _, str := range strings {
		set[str] = true
	}
	return set
}

func main() {
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}
	mySet := createSet(sequence)
	fmt.Println(mySet)
}
