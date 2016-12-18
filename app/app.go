package app

import (
	"fmt"
	"net/http"
	"html/template"
)

var _ = fmt.Print

func Index(w http.ResponseWriter, r *http.Request) {
	html := `
<html>
<head>
<title>go-e2e-test-sample</title>
</head>
<body>
<div>
<p>Hello {{ .Name }}</p>
<form method="post" action="/">
<input type="text" name="name" value="">
<input type="submit" value="send">
</form>
</div>
</body>
</html>
	`

	t, err := template.New("index").Parse(html)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	type Args struct {
		Name string
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	name := r.FormValue("name")
	if err := t.Execute(w, &Args{Name: name}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
