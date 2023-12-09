package designer

import "flag"

type Designer struct {
	ProjectName string

	Layers []Layer
}

type Layer struct {
	Name      string
	Path      string
	Blueprint Blueprint
}

type Blueprint interface {
	Generate() error
	SetFlags(...flag.Flag) error
	SetFlagsFromMap(map[string]string) error
}
