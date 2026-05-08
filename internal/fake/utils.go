package fake

import "fmt"

type RunTimeDocsMethod struct {
	Name        string
	Description string
	Params      []string
}

type RunTimeDocs struct {
	Struct  string
	Methods map[string]RunTimeDocsMethod
}

var ErrLocaleNotSupported = fmt.Errorf("locale not supported")
