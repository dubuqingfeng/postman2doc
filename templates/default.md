# {{.Name}}
{{range .Requests}}
## {{.Name}}
{{if .Description}}
{{.Description}}
{{end}}
{{if .URL}}
`{{.Method}} {{.URL}}`
{{end}}
### 参数
{{if .PayloadRaw}}
```
{{.PayloadRaw}}
```
{{else}}
无
{{end}}
{{if .Responses}}
### 返回值
```
{{range .Responses}}
{{.Body}}
{{end}}
```
{{end}}
{{end}}