package radarapi

import "strings"

func fieldsToCommaString(fields []string) string {
	var builder strings.Builder
	for i, f := range fields {
		builder.WriteString(f)
		if i < len(fields)-1 {
			builder.WriteString(",")
		}
	}
	return builder.String()
}
