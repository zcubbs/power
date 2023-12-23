package apiserver

import _ "embed"

//go:embed files/main.go.tmpl
var MainFileTemplate string
