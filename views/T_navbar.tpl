{{define "navbar"}}
<a class="navbar-brand" href="/"> My blog</a>

<nav>
  <ul class="nav navbar-nav">
    <li {{if .IsHome}}class="active" {{end}}><a href="/">Home</a></li>
    <li {{if .IsCategory}}class="active" {{end}}><a href="/category">Category</a></li>
    <li {{if .IsTopic}}class="active" {{end}}><a href="/topic">Topic</a></li>
  </ul>
</nav>

<nav class="pull-right">
  <ul class="nav navbar-nav">
    {{if .IsLogin}}
    <li><a href="/login?logout=true">Logout</a></li>
    {{else}}
    <li><a href="/login">Login</a></li>
    {{end}}
  </ul>
</nav>
{{end}}
