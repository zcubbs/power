package main

import (
	_ "embed"
	"github.com/zcubbs/power/pkg/blueprint"
)

// Plugin is the exported plugin blueprint.
var Plugin blueprint.Generator = &Generator{}
