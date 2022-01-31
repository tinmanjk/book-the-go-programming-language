package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Structs")
	var dilbert Employee
	fmt.Println(dilbert)
	dilbert.Salary -= 5000 // demoted, for writing too few lines of code
	position := &dilbert.Position
	*position = "Senior " + *position // promoted, for outsourcing to Elbonia
	fmt.Println(dilbert.Position)     // "Senior"

	// dot notation with pointer to a struct
	var employeeOfTheMonth *Employee = &dilbert               // pointer struct variable
	employeeOfTheMonth.Position += " (proactive team player)" // implicit dereferencing
	fmt.Println(employeeOfTheMonth.Position)                  // Senior  (proactive team player)
	(*employeeOfTheMonth).Position += " explicit"             // explicit dereferencing
	fmt.Println(employeeOfTheMonth.Position)                  // Senior  (proactive team player explicit)

	// return from function
	fmt.Println(EmployeeByID(dilbert.ManagerID).Position) // "Pointy-haired boss"

	id := dilbert.ID
	EmployeeByID(id).Salary = 0 // fired for... no	real reason
	// EmployeeByIDNoPointer(id).Salary = 0 // compilation error: no variable -> need to assign to variable first

	arrToSort := []int{3, 2, 4, 1}
	Sort(arrToSort)        // in place
	fmt.Println(arrToSort) // [1 2 3 4]

	fmt.Println("\nEmpty Struct Literal") // cumbersome syntax -> to be avoided
	emptyStruct := struct{}{}             // the type is struct{} -> mind the double {}{}
	fmt.Println(emptyStruct)
	// used instead of bool for hashsets
	seen := make(map[string]struct{}) // set of strings
	if _, ok := seen["non-existent"]; !ok {
		seen["non-existent"] = struct{}{} // ...first time seeing s...
		fmt.Println(seen["non-existent"])
	}

	fmt.Println("\n4.4.1 Struct Literals")
	// first form - fragile, well known structs inside its own package etc
	p := Point{1, 2} // if no field name - EVERY FIELD in CORRECT ORDER
	fmt.Println(p)   // {1 2}
	// second form - prefer overall
	p = Point{Y: 2} // X is set to zero-value
	// p = Point{2}    // would not compile
	fmt.Println(p) // {0 2}

	fmt.Println("\nPassing and returning structs to/from functions")
	fmt.Println(Scale(Point{1, 2}, 5)) // "{5 10}"

	emp := dilbert
	emp.Salary = 100
	fmt.Println(Bonus(&emp, 50)) // 50 -> indirect passing
	AwardAnnualRaise(&emp)
	fmt.Println(emp.Salary) // 105

	pp := &Point{1, 2} // short notation for creation of pointer to struct
	// can be used within expression
	fmt.Println(*pp, &Point{3, 4}) // {1 2} &{3 4}
}

type Point struct{ X, Y int }

// arguments and returned values from functions
func Scale(p Point, factor int) Point {
	return Point{p.X * factor, p.Y * factor}
}

func Bonus(e *Employee, percent int) int {
	return e.Salary * percent / 100
}

func AwardAnnualRaise(e *Employee) {
	e.Salary = e.Salary * 105 / 100
}

// Field order IS significant
type Employee struct {
	ID            int
	Name, Address string // combining fields of the same type - only related not of the same type
	//lint:ignore U1000 unused
	doB       time.Time // not exported
	Position  string
	Salary    int
	ManagerID int
}

func EmployeeByID(id int) *Employee {
	return &Employee{Position: "Pointy-haired boss"}
}

func EmployeeByIDNoPointer(id int) Employee {
	return Employee{Position: "Pointy-haired boss"}
}

// binarty tree for insertion sort
type tree struct {
	value       int
	left, right *tree
	// cycle       tree // cycle declaration not allowed
}

// Sort sorts values in place.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root) // smart way to reuse the underlying array with append
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left) // only goal is to go to the bottom of the left in the recursion
		values = append(values, t.value)      // real work
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}
