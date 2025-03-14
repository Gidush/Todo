package validation

import (
	"errors"
	"fmt"
	"strings"
)

type ResultBuilder struct {
	results []string
}

func (v *ResultBuilder) Add(isValid bool, message string, a ...any) {
	if !isValid {
		v.results = append(v.results, fmt.Sprintf(message, a...))
	}
}

func (v *ResultBuilder) String() string {
	if len(v.results) == 0 {
		return ""
	}
	return strings.Join(v.results, "\n")
}

func (v *ResultBuilder) Validate() error {
	if len(v.results) == 0 {
		return nil
	}
	return errors.New(strings.Join(v.results, "\n"))
}
