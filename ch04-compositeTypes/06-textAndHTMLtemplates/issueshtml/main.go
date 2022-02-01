// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package main

import (
	"log"
	"os"

	htmlTemplate "html/template"
	textTemplate "text/template"

	"github.com/tinmanjk/tgpl/ch04-compositeTypes/05-json/github"
)

var template = `
<h1>{{.TemplateVersion}} - {{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
</tr>
{{range .Items}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`

// go run . > issuesHTMLvsText.html
func main() {

	// issues for escape demonstration
	result, err := github.SearchIssues([]string{"repo:golang/go 3133 10535"})
	if err != nil {
		log.Fatal(err)
	}
	var issueListText = textTemplate.Must(textTemplate.New("issuelistText").Parse(template))
	result.TemplateVersion = "Text Template"
	// text -> does not escape &lt; in string to be taken literally
	// instead passes it to the HTML renderer - so it becomes '<'
	if err := issueListText.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}

	var issueListHTML = htmlTemplate.Must(htmlTemplate.New("issuelistHTML").Parse(template))
	result.TemplateVersion = "HTML Template"
	// **auto-escaping behavior with html template** -> purpose the **input** should NOT alter the HTML structure
	// e.g. input having a script that executes <script etc> -> this will be escaped and rendered as text
	// &lt; ('<') will be DOUBLE-escaped, so it's not treated as it escapes '<' but as a literal &lt;
	if err := issueListHTML.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}

	// turn off auto-escaping by using template.HTML (.CSS, JS) types
	const templ = `<p>StringField: {{.StringField}}</p>
<p>HTMLField: {{.HTMLField}} </p>`
	t := htmlTemplate.Must(htmlTemplate.New("escape").Parse(templ))
	var data struct {
		StringField string            // untrusted plain text
		HTMLField   htmlTemplate.HTML // trusted HTML
	}
	data.StringField = "<b>Hello!</b>"
	data.HTMLField = "<b>Hello!</b>"
	if err := t.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}

}
