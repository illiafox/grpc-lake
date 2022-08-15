package errors

import (
	"strings"
)

func (i InternalError) Wrap(scope string) error {

	if i.Scope != "" {
		var builder strings.Builder
		builder.Grow(len(i.Scope) + len(Separator) + len(scope))

		builder.WriteString(i.Scope)
		builder.WriteString(Separator)
		builder.WriteString(scope)

		i.Scope = builder.String()
	} else {
		i.Scope = scope
	}

	return i
}
