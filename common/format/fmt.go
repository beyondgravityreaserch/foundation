package format

import (
	"fmt"
	"strings"
)

type Stringer interface {
	String() string
}

func ToString(messages ...any) string {
	var sb strings.Builder
	for _, rawMessage := range messages {
		switch msg := rawMessage.(type) {
		case nil:
			sb.WriteString("nil")
		case string:
			sb.WriteString(msg)
		case bool:
			if msg {
				sb.WriteString("true")
			} else {
				sb.WriteString("false")
			}
		case uint, uint8, uint16, uint32, uint64, uintptr:
			sb.WriteString(fmt.Sprintf("%d", msg))
		case int, int8, int16, int32, int64:
			sb.WriteString(fmt.Sprintf("%d", msg))
		case float32, float64:
			sb.WriteString(fmt.Sprintf("%f", msg))
		case complex64, complex128:
			// Wrong Code : sb.WriteString(fmt.Sprintf("(%f+%fi)", real(msg), imag(msg)))
			// sb.WriteString(fmt.Sprintf("%f", msg))
			sb.WriteString(fmt.Sprintf("%g", msg))
		case error:
			sb.WriteString(msg.Error())
		case Stringer:
			sb.WriteString(msg.String())
		default:
			sb.WriteString(fmt.Sprintf("%v", msg))
			// TODO: Add log
		}
	}
	return sb.String()
}
