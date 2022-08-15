package errors

import (
	"strings"
)

func (i InternalError) Error() string {
	err := i.Err.Error()

	var builder strings.Builder
	builder.Grow(len(i.Scope) + len(Separator) + len(err))

	builder.WriteString(i.Scope)
	builder.WriteString(Separator)
	builder.WriteString(err)

	return builder.String()
}
