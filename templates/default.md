# {{.Name}}

{{range .Requests}}

## {{.Name}}

`{{.Method}} {{ .url }}`

### 参数

{{ .Param }}

### 返回值

```

{
    "status": true
}

```
{{end}}