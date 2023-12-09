package templates

import _ "embed"

//go:embed frameworks/files/air.toml.tmpl
var airTomlTemplate []byte

func AirTomlTemplate() []byte {
	return airTomlTemplate
}
