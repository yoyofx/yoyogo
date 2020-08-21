<html>
<head>
<title>{{.Title}}</title>
{{template "baseStyle" .}}
</head>

<body>
<div class="redTitle">CSRF_TOKEN:{{.CSRF_TOKEN}}</div>
<div>{{.shareIndex}}</div>
{{.Desc|unescaped}}

{{if ge  .Status .OpStatus}}
123
{{else}}
456
{{end}}

<ul>

{{range $i,$v:=.ItemArr}}
<li>list item {{$i}}=>{{$v}}</li>
{{end}}
</ul>

{{.orange}}
</body>
</html>