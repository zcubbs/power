package main

import (
	_ "embed"
	"github.com/zcubbs/power/pkg/blueprint"
)

var Plugin blueprint.Generator = &Generator{}
