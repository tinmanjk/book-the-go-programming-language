package main

import (
	"fmt"
	"log"
	"os"
	textTemplate "text/template"
	"time"

	"github.com/tinmanjk/tgpl/ch04-compositeTypes/05-json/github"
)

func main() {
	fmt.Println("Text Templating")
	// fist step - parse the template into suitable structure -> template.Template
	var report = textTemplate.Must(textTemplate.New("issuelist").
		Funcs(textTemplate.FuncMap{"daysAgo": daysAgo}).
		Parse(templ)) // Must is a pattern that generates a panic if error in the chain, otherwise report

	take := 2
	result, err := github.SearchIssues([]string{"repo:golang/go", "is:open", "json", "decoder"})
	if err != nil {
		log.Fatal(err)
	}
	result.Items = result.Items[:take]
	result.TotalCount = take
	// second step -> execute on specific inputs -> github.IssuesSearchResult
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}

	// see issueshtml for html templating (vs text templating)
}

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

// . is current value
// . -> initialy as the template parameter -> github.IssuesSearchResult
// {{range .Items}} and {{end}} -> create a loop -> . bound to elements of Items
// | -> like Unix-piping
// printf -> fmt.Sprintf
const templ = `{{.TotalCount}} issues:
{{range .Items}}----------------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}`
