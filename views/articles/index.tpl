{{range $key, $article := .Articles}}
    <p>{{$key}}</p>
    <p>{{$article.Title}}</p>
    <p>{{$article.Description}}</p>
    <p>{{$article.Date}}</p>
    <a href="/articles/{{$article.Id}}">Show</a>
    <a href="/articles/{{$article.Id}}/delete">Delete</a>
    <a href="/articles/{{$article.Id}}/edit">Edit</a>
{{end}}
