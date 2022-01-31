package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Maps")

	// with built-in make function
	ages := make(map[string]int) // mapping from strings to ints
	fmt.Println(ages)
	// with map literal
	ages = map[string]int{
		"alice":   31,
		"charlie": 34,
	}
	fmt.Println(ages)
	ages["alice"] = 32
	fmt.Println(ages["alice"]) // "32"

	// removed via delete built-in function
	delete(ages, "alice") // remove element ages["alice"]
	fmt.Println(ages)

	// map lookup that doen't find anything -> ZERO VALUE
	fmt.Println(ages["bob"]) // "0"
	ages["bob"] += 1
	ages["bob"]++
	fmt.Println(ages["bob"]) // "2"

	// _ = &ages["bob"] // compile error: cannot take address of map element
	// reason growing a map - restoring in other locations

	// iteration
	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}

	// sorting
	fmt.Println("\nSorting")
	names := make([]string, 0, len(ages)) // since we know the capacity from the beginning
	// len is set to 0 cause otherwise append **will start at len not 0**!!!
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}

	fmt.Println("\nZero Value:")
	{
		var ages map[string]int
		fmt.Println(ages == nil)    // "true"
		fmt.Println(len(ages) == 0) // "true"

		delete(ages, "non-existent")
		fmt.Println(ages["non-existent"]) // 0
		for _, v := range ages {          // would print nothing
			fmt.Println(v)
		}
		// **however** assignment -> panic - also static check - SA5000
		// ages["carol"] = 21 // panic: assignment to entry in nil map
		// The map MUST be allocated before it can store something in itself

	}

	fmt.Println("\nValue, Ok Test")
	if age, ok := ages["bob"]; ok {
		fmt.Println(age)
	}

	fmt.Println("\nComparisons")
	// only with nil ==
	fmt.Println(ages == nil)       // false
	fmt.Println(equal(ages, ages)) // true

	fmt.Println("\nMaps as Sets")
	set := map[string]bool{}
	set["existing"] = true
	fmt.Println(set["existing"])
	fmt.Println(set["non-existent"])

	fmt.Println("\nSlices as Keys") // see

	var m = make(map[string]int) // actually slice of strings as key but here string with k function helper
	// concatenation of all the strings in the slice
	k := func(list []string) string {
		return fmt.Sprintf("%q", list) // -> %q '' for **string boundaries**
	}

	Add := func(list []string) {
		m[k(list)]++
	}

	Count := func(list []string) int {
		return m[k(list)]
	}

	sliceStrings := []string{"one", "two"}
	Add(sliceStrings)
	sliceStrings = append(sliceStrings, "three")
	Add(sliceStrings)
	fmt.Println(Count(sliceStrings))   // 1
	diffSlice := make([]string, 0, 50) // remember 0 length
	diffSlice = append(diffSlice, "one")
	diffSlice = append(diffSlice, "two")
	Add(diffSlice)
	fmt.Println(Count(diffSlice)) // 2

	fmt.Println("\nComposite value types")
	var graph = make(map[string]map[string]bool) // [node]hashset of edges

	// idiomatic way of populating a map LAZILY
	addEdge := func(from, to string) {
		edges := graph[from]
		if edges == nil { // lazy initiation
			edges = make(map[string]bool)
			graph[from] = edges
		}
		edges[to] = true
	}

	hasEdge := func(from, to string) bool {
		// **kind of chaining of zero-value return of lookup in nil map**
		// graph[from] = nil -> nil[to] doesn't make sense but it knows to return its zero value = false
		return graph[from][to]
	}

	addEdge("first", "second")
	fmt.Println(hasEdge("first", "second")) // true
	fmt.Println(graph["second"] == nil)
	fmt.Println(graph["second"]["first"]) // this would be null-reference exception in other languages
	// however it still returns the zero value for the second map -> false
	// zero-value IS **NOT STORED** OR INITIALIZED like in Arrays -> it is simply RETURNED!!!
	// https://stackoverflow.com/questions/54124439/map-initialization-in-go

}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}
