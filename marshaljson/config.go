package marshaljson

import "reflect"

const (
	tabDefault       = "default"
	tabDefaultString = "defaultString"
	tabDateTime      = "datetime"
)

type tabT struct {
	refTypOf reflect.Type
	restrain string
	fun      typeConvT
}

var (
	tabList = []string{tabDefault, tabDefaultString, tabDateTime}
	tabMap  = map[string]tabT{
		tabDefault: {
			refTypOf: reflect.TypeOf(defaultT{}),
			restrain: "",
			fun:      defaultT{},
		},
		tabDefaultString: {
			refTypOf: reflect.TypeOf(defaultStringT{}),
			restrain: "",
			fun:      defaultStringT{},
		},
		tabDateTime: {
			refTypOf: reflect.TypeOf(dateTime{}),
			restrain: "time.Time",
			fun:      dateTime{},
		},
	}
)
