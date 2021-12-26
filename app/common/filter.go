package common

import "github.com/importcjj/sensitive"

func Sfilter() *sensitive.Filter {
	filter := sensitive.New()
	return filter
}
