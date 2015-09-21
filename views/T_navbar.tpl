{{define "navbar"}}
<a class="navbar-brand" href="/"> {{i18n .Lang .title_lang}}</a>

<nav>
  <ul class="nav navbar-nav">
    <li {{if .IsHome}}class="active" {{end}}><a href="/">{{i18n .Lang .home_lang}}</a></li>
    <li {{if .IsCategory}}class="active" {{end}}><a href="/category">{{i18n .Lang .cate_lang}}</a></li>
    <li {{if .IsTopic}}class="active" {{end}}><a href="/topic">{{i18n .Lang .topic_lang}}</a></li>
  </ul>
</nav>

<nav class="pull-right">
  <ul class="nav navbar-nav">
    {{if .IsLogin}}
    <li><a href="/login?logout=true">{{i18n .Lang .logout_lang}}</a></li>
    {{else}}
    <li><a href="/login">{{i18n .Lang .login_lang}}</a></li>
    {{end}}
</nav>
<div class="collapse navbar-collapse" id="bs-example-navbar-collapse-1">
    <ul class="nav navbar-nav navbar-right">
      <li class="dropdown">
        <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">Language <span class="caret"></span></a>
        <ul class="dropdown-menu">
          <li><a href="/?lang=en-US">English</a></li>
          <li><a href="/?lang=zh-HK">正體中文</a></li>
          <li><a href="/?lang=zh-CN">简体中文</a></li>
        </ul>
      </li>
    </ul>
  </div><!-- /.navbar-collapse -->
{{end}}
