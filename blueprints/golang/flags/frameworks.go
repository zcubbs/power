package flags

import (
	"fmt"
	"strings"
)

type Framework string

const (
	Chi             Framework = "chi"
	Gin             Framework = "gin"
	Fiber           Framework = "fiber"
	GorillaMux      Framework = "gorilla/mux"
	HttpRouter      Framework = "httprouter"
	StandardLibrary Framework = "standard-library"
)

var AllowedProjectTypes = []string{
	string(Chi),
	string(Gin),
	string(Fiber),
	string(GorillaMux),
	string(HttpRouter),
	string(StandardLibrary),
}

func (f *Framework) String() string {
	return string(*f)
}

func (f *Framework) Type() string {
	return "Framework"
}

func (f *Framework) Set(value string) error {
	for _, project := range AllowedProjectTypes {
		if project == value {
			*f = Framework(value)
			return nil
		}
	}

	return fmt.Errorf("framework to use. Allowed values: %s", strings.Join(AllowedProjectTypes, ", "))
}
