package main

import (
	"database/sql"
	"fmt"
)

func main() {
	fmt.Println("Type Switches")
}

//lint:ignore U1000 ...
func listTracks(db sql.DB, artist string, minYear, maxYear int) {
	// Go API -> distinguish between fixed and variable parts( ? )
	result, err := db.Exec(
		"SELECT * FROM tracks WHERE artist = ? AND ? <= year AND year <= ?",
		artist, minYear, maxYear)
	fmt.Println(result, err)
}

// type assertion style without type switch statement
//lint:ignore U1000 ...
func sqlQuote(x interface{}) string {
	if x == nil {
		return "NULL"
	} else if _, ok := x.(int); ok {
		return fmt.Sprintf("%d", x)
	} else if _, ok := x.(uint); ok {
		return fmt.Sprintf("%d", x)
	} else if b, ok := x.(bool); ok {
		if b {
			return "TRUE"
		}
		return "FALSE"
	} else if s, ok := x.(string); ok {
		return s
		// return sqlQuoteString(s) // (not shown)
	} else {
		panic(fmt.Sprintf("unexpected type %T: %v", x, x))
	}
}

//lint:ignore U1000 ...
func sqlQuoteWithTypeSwitch(x interface{}) string {
	// Although the type of x is interface{}, we consider it a
	// discriminated union of int, uint, bool, string, and nil.
	// -> panic otherwise

	// -> type switch evaluates the expression x.(type) -> type is keyword
	// AND assigns to x (special form allowing assignment vs normal switch expression)
	switch x := x.(type) { // assignment(shadowing) to avoid type assertions in case statements
	case nil: // -> nil interface value, dynamic type is nil
		// https://go.dev/ref/spec#Type_switches
		// that case is used when the expression in the TypeSwitchGuard is a nil interface value.
		return "NULL"
	case int, uint:
		return fmt.Sprintf("%d", x) // !!! here x is still interface{} because multiple types in case
	case bool:
		if x {
			return "TRUE"
		}
		// fallthrough -> not possible in type switch
		return "FALSE"
	case string:
		return x
	default:
		panic(fmt.Sprintf("unexpected type %T: %v", x, x))

	}
}
