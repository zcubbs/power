package buildins_helloworld

import _ "embed"

//go:embed files/hello.txt.tmpl
var helloTxtTemplate string
