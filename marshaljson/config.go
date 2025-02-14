package marshaljson

import "reflect"

const (
	tabNameDefault       = "default"
	tabNameDefaultString = "defaultString"
	tabNameDateTime      = "datetime"
)

type tabT struct {
	refTypOf reflect.Type
	restrain string
	fun      typeConvT
}

var (
	tabList = []string{tabNameDefault, tabNameDefaultString, tabNameDateTime}
	tabMap  = map[string]tabT{
		tabNameDefault: {
			refTypOf: reflect.TypeOf(tabDefault{}),
			restrain: "",
			fun:      tabDefault{},
		},
		tabNameDefaultString: {
			refTypOf: reflect.TypeOf(tabDefaultString{}),
			restrain: "",
			fun:      tabDefaultString{},
		},
		tabNameDateTime: {
			refTypOf: reflect.TypeOf(tabDateTime{}),
			restrain: "time.Time",
			fun:      tabDateTime{},
		},
	}
)
