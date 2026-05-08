package fake

import "fmt"

type RunTimeDocs struct {
	Domain      string
	Method      string
	Description string
	Params      []string
}

var ErrLocaleNotSupported = fmt.Errorf("locale not supported")
