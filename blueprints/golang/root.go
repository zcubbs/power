package golang

import "flag"

type Blueprint struct{}

func (b *Blueprint) Generate() error {
	return nil
}

func (b *Blueprint) SetFlags(...flag.Flag) error {
	return nil
}

func (b *Blueprint) SetFlagsFromMap(map[string]string) error {
	return nil
}
