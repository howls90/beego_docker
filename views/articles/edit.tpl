<form method="POST" action="/articles/{{.Article.Id}}/edit" method="post">
    <label> Title </label>
    <input type="text" name="title" value="{{.Article.Title}}"><br />
    <label> Description </label>
    <input type="text" name="description" value="{{.Article.Description}}"/><br />
    <input type="submit" value="Save" />
</form>
