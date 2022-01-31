package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/tinmanjk/tgpl/ch04-compositeTypes/05-json/github"
)

func main() {

	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false,
			Actors: []string{"Humphrey Bogart",
				"Ingrid Bergman"}, unexported: "unexported"},
		{Title: "Cool Hand Luke", Year: 1967, Color: true,
			Actors: []string{"Paul Newman"}, unexported: "unexported"},
		{Title: "Bullitt", Year: 1968, Color: true,
			Actors: []string{"Steve McQueen",
				"Jacqueline Bisset"}, unexported: "unexported"},
	}
	fmt.Printf("Go slice of movies -> %#v\n", movies)

	// marshaling -> Go->Json conversion
	fmt.Println("\nMarshalling Go->JSON conversion")
	data, err := json.Marshal(movies) // uses reflection to determine the field names
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("\n%s\n", data) //%s for byte slices decodes UTF8
	// [{"Title":"Casablanca","released":1942,"Actors":["Humphrey Bogart","Ingrid Bergman"]},
	// {"Title":"Cool Hand Luke","released":1967,"color":true,"Actors":["Paul Newman"]},
	// {"Title":"Bullitt","released":1968,"color":true,"Actors":["Steve McQueen","Jacqueline Bisset"]}]
	// -> Year = released, Color = color BUT if false -> it's omitted see Casablanca
	data, _ = json.MarshalIndent(movies, "", "  ") // human readable
	fmt.Printf("%s\n", data)

	fmt.Println("\nUnmarshalling Go->JSON conversion")
	var titles []struct{ Title string }                   // anonymous struct
	if err := json.Unmarshal(data, &titles); err != nil { // needs pointer to storage location
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Println(titles) // "[{Casablanca} {Cool Hand Luke} {Bullitt}]"

	// Issues
	{
		result, err := github.SearchIssues([]string{"repo:golang/go", "is:open", "json", "decoder"})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d issues:\n", result.TotalCount)
		for _, item := range result.Items {
			fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
		}
	}
}

type Movie struct {
	Title string
	Year  int  `json:"released"`        // field tags
	Color bool `json:"color,omitempty"` // field tags
	// by convention field tags: space-separated list of key:"value" pairs;
	// json key behavior for -> encoding/json package
	// alternative name: "name, option"
	// omitempty -> if zero-value or otherwise empty do not output
	Actors     []string
	unexported string // will not be marshalled
}
