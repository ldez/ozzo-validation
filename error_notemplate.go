//go:build ozzo_validation_no_template
// +build ozzo_validation_no_template

package validation

import (
	"fmt"
	"strings"
)

// Error returns the error message.
func (e ErrorObject) Error() string {
	if len(e.params) == 0 {
		return e.message
	}

	for k, v := range e.params {
		e.message = strings.ReplaceAll(e.message, "{{."+k+"}}", fmt.Sprintf("%v", v))
	}

	return e.message
}
