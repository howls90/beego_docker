{{range $key, $article := .Articles}}
    <p>{{$key}}</p>
    <p>{{$article.Title}}</p>
    <p>{{$article.Description}}</p>
    <p>{{$article.Date}}</p>
{{end}}