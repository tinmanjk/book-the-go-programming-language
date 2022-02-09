package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

func main() {
	fmt.Println("Sorting with sort.Slice")

	fmt.Println("\nMine: ")
	{
		fmt.Println("\nSorting integers")
		// main API
		// 1. Sort - unstable sort - sort.Ints convenience
		// 2. Stable - stable sort [original order of equal elements] - NO sort.StableInts
		// 3. IsSorted - sort.IntsAreSorted convenience
		// 4. Reverse -> no sort.ReverseInts
		// -> Type convenience
		// 5. sort.Ints, sort.Strings, sort Float64s
		// -> Slice convenience
		// sort.Slice -> []interface{} and a less function

		is := []int{3, 2, 1}
		// sort.Sort(is)        // does not work -> needs sort.Interface which []int does not satisfy
		fmt.Println("\nManual Conversion to sort.IntSlice")
		sliceConverted := sort.IntSlice(is) // named typed conversion -> copying len/cap -> same underlying array
		sliceConverted.Sort()
		fmt.Println(is)             // [1 2 3]
		fmt.Println(sliceConverted) // [1 2 3] - obv same underlying arary

		// unsorting
		is[0], is[2] = is[2], is[0]
		fmt.Println(is) // [3 2 1]

		fmt.Println("\nsort.Ints convenience")
		sort.Ints(is)   // just wrapper around conversion to sort.IntSlice and callin Sort there
		fmt.Println(is) // [1 2 3]

		fmt.Println("\nStable sort - need sort.Interface")
		// unsorting
		is[0], is[2] = is[2], is[0]
		fmt.Println(is) // [3 2 1]

		sort.Stable(sliceConverted) // no conveience sort.IntsStable
		fmt.Println(is)             // [1 2 3]
		// same mechanic for Strings and Floats

		fmt.Println("\nReverse functionality")
		toBeReversed := sort.Reverse(sliceConverted) // JUST flips the Less Function, NO SORT
		fmt.Println(toBeReversed)                    // &{[1 2 3]} -> still the same data
		sort.Sort(toBeReversed)                      // -> NOW sorts in reverse manner
		fmt.Println(sliceConverted)                  //[3 2 1]
		fmt.Println(toBeReversed)                    //&{[3 2 1]}

		sort.Sort(sort.Reverse(sliceConverted)) // one-liner
		fmt.Println(sliceConverted)             // still [3 2 1] - they are SORTED in reverse manner
		sort.Sort(sliceConverted)
		fmt.Println(sliceConverted) //  [1 2 3]
		sort.Sort(sort.Reverse(sliceConverted))
		fmt.Println(sliceConverted) //  [3 2 1]
	}
	fmt.Println("\n---\\Mine")

	fmt.Println("\nImplementation for string slice")
	names := []string{"cd", "ab", "ef"}
	// sort.Sort(StringSlice(names)) // one-liner from book - **not expensive** - two converions to named type - interface value
	// see below:

	fmt.Println("\nMine: Decomposition of conversions from []string to StringSlice to sort.Interface")
	{
		// ** MINE -> Decomposition **
		// main.go:66      0x4a53fa        48899424f8010000                mov qword ptr [rsp+0x1f8], rdx
		// main.go:66      0x4a5402        48c784240002000003000000        mov qword ptr [rsp+0x200], 0x3
		// main.go:66      0x4a540e        48c784240802000003000000        mov qword ptr [rsp+0x208], 0x3
		conversionToSatisfyingType := StringSlice(names) // simple copy of the fields of the slice struct
		// main.go:73      0x4a541a        4889942428020000                mov qword ptr [rsp+0x228], rdx
		// main.go:73      0x4a5422        48c784243002000003000000        mov qword ptr [rsp+0x230], 0x3
		// main.go:73      0x4a542e        48c784243802000003000000        mov qword ptr [rsp+0x238], 0x3
		conversionToInterface := sort.Interface(conversionToSatisfyingType) // $runtime.convTslice is called for the interface conversion
		// main.go:77      0x4a543a        488b842428020000                mov rax, qword ptr [rsp+0x228]
		// main.go:77      0x4a5442        bb03000000                      mov ebx, 0x3
		// main.go:77      0x4a5447        4889d9                          mov rcx, rbx
		// main.go:77      0x4a544a        e8114ef6ff                      call $runtime.convTslice (rax pointer to array, rbx, rcx len cap)
		// func convTslice(val []byte) (x unsafe.Pointer) {
		// 	// Note: this must work for any element type, not just byte.
		// 	if (*slice)(unsafe.Pointer(&val)).array == nil {
		// 		x = unsafe.Pointer(&zeroVal[0])
		// 	} else {
		// 		x = mallocgc(unsafe.Sizeof(val), sliceType, true)
		// 		*(*[]byte)(x) = val -> copying the fields of the slice -> **INCLUDING THE ADDRESS OF THE UNDERLYING ARRAY**
		// 	}
		// 	return
		// }
		// main.go:77      0x4a544f        4889442440                      mov qword ptr [rsp+0x40], rax -> returned pointer to slice struct (three fields)
		// main.go:77      0x4a5454        488d15e5850200                  lea rdx, ptr [rip+0x285e5] -> type descriptor
		// main.go:77      0x4a545b        4889942490010000                mov qword ptr [rsp+0x190], rdx
		// main.go:77      0x4a5463        4889842498010000                mov qword ptr [rsp+0x198], rax

		sort.Sort(conversionToInterface)
		// main.go:88      0x4a546b        4889c3                          mov rbx, rax
		// main.go:88      0x4a546e        4889d0                          mov rax, rdx
		// main.go:88      0x4a5471        e8ea2dfeff                      call $sort.Sort // rax = type descriptor, rbx = value (pointer to allocated slice)
	}

	fmt.Println(names) // [ab cd ef]
	// sort.StringSlice available
	// unsort
	names[0], names[2] = names[2], names[0]
	fmt.Println(names)

	sort.Strings(names) // conevenience
	fmt.Println(names)

	var tracks = []*Track{
		{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
		{"Go", "Moby", "Moby", 1992, length("3m37s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	}

	printTracks(tracks)
	sort.Sort(byArtist(tracks))
	printTracks(tracks)

	fmt.Println("\nReverse Artist")
	// uses composition -> reverse struct with Interface embedding
	// method promotion of sort.Interface but **Less is overridden**

	// type reverse struct{ Interface } // that is sort.Interface
	// func (r reverse) Less(i, j int) bool { return r.Interface.Less(j, i) }
	// func Reverse(data Interface) Interface { return reverse{data} }

	sort.Sort(sort.Reverse(byArtist(tracks))) // sort.Reverse return sort.Interface with reversed Less function
	printTracks(tracks)

	fmt.Println("\nSort by Year")
	sort.Sort(byYear(tracks))
	printTracks(tracks)

	fmt.Println("\nCustom Sort")
	byAlbum := func(x, y *Track) bool {
		return x.Album < y.Album
	}
	custSort := customSort{tracks, byAlbum}
	sort.Sort(custSort)
	printTracks(tracks)

	fmt.Println("\nMulti-tier ordering function")
	titleYearLengthFunc := func(x, y *Track) bool {
		if x.Title != y.Title {
			return x.Title < y.Title
		}

		if x.Year != y.Year {
			return x.Year < y.Year
		}

		return x.Length < y.Length
	}
	titleYearLength := customSort{tracks, titleYearLengthFunc}
	sort.Sort(titleYearLength)
	printTracks(tracks)

	fmt.Println("\nIs Sorted n-1 complexity")
	values := []int{3, 1, 4, 1}
	fmt.Println(sort.IntsAreSorted(values)) // "false"
	sort.Ints(values)
	fmt.Println(values)                     // "[1 1 3 4]"
	fmt.Println(sort.IntsAreSorted(values)) // "true"
	sort.Sort(sort.Reverse(sort.IntSlice(values)))
	fmt.Println(values)                     // "[4 3 1 1]"
	fmt.Println(sort.IntsAreSorted(values)) // "false"
}

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

// tabwriter package interesting
func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
	fmt.Println()
}

type byArtist []*Track // named type for sort.Interface satisfaction

func (x byArtist) Len() int           { return len(x) }
func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] } // faster swap due to pointers

type byYear []*Track

func (x byYear) Len() int           { return len(x) }
func (x byYear) Less(i, j int) bool { return x[i].Year < x[j].Year }
func (x byYear) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (c customSort) Len() int           { return len(c.t) }
func (c customSort) Less(i, j int) bool { return c.less(c.t[i], c.t[j]) }
func (c customSort) Swap(i, j int)      { c.t[i], c.t[j] = c.t[j], c.t[i] }

// from package sort
type Interface interface {
	Len() int
	Less(i, j int) bool // function for comparing two elements, accessed by integer identifiers
	Swap(i, j int)
}

// satisfaction of Interface by IntSlice from sort.Package
// addidtional FloatSlice and StringSlice
type IntSlice []int

func (x IntSlice) Len() int           { return len(x) }
func (x IntSlice) Less(i, j int) bool { return x[i] < x[j] }
func (x IntSlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

// Ints sorts a slice of ints in increasing order.
// func Ints(x []int) { Sort(IntSlice(x)) }

type StringSlice []string

func (p StringSlice) Len() int {
	return len(p)
}
func (p StringSlice) Less(i, j int) bool {
	return p[i] < p[j]
}
func (p StringSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
