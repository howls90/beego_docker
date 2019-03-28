<form method="POST" action="/articles/save" method="post">
    {{ .xsrfdata }}
    <label> Title </label>
    <input type="text" name="title"><br />
    <label> Description </label>
    <input type="text" name="description"/><br />
    <input type="submit" value="Save" />
</form>
