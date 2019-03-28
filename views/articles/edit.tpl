<form method="POST" action="/articles/{{.Article.Id}}/edit" method="post">
    {{ .xsrfdata }}
    <input type="text" name="id" value="{{.Article.Id}}" hidden><br />
    <label> Title </label>
    <input type="text" name="title" value="{{.Article.Title}}"><br />
    <label> Description </label>
    <input type="text" name="description" value="{{.Article.Description}}"/><br />
    <input type="submit" value="Save" />
</form>
