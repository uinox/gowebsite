{{define "navbar"}}
<a href="/" class="navbar-brand">Home</a>
<div>
    <ul class="nav navbar-nav">
        <li {{if .IsHome}}class="active"{{end}}><a href="/">Home</a></li>
        <li {{if .IsCategory}}class="active"{{end}}><a href="/category">Category</a></li>
        <li {{if .IsTopic}}class="active"{{end}}><a href="/topic">Article</a></li>
    </ul>
</div>

<div class="pull-right">
    <ul class="nav navbar-nav">
        {{if .IsLogin}}
        <li><a href="/login?exit=true">Exit</a></li>
        {{else}}
        <li><a href="/login">Login</a></li>
        {{end}}
    </ul>
</div>
{{end}}