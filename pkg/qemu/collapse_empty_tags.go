package qemu

import (
	"fmt"
	"regexp"
	"strings"
)

func CollapseEmptyTags(input string) string {
	r := regexp.MustCompile(`<([a-zA-Z0-9]+)(.*?)>(.*?)</([a-zA-Z0-9]+)>`)

	return r.ReplaceAllStringFunc(
		input,
		func(match string) string {
			tag := r.ReplaceAllString(match, "$1")
			attributes := r.ReplaceAllString(match, "$2")
			content := r.ReplaceAllString(match, "$3")
			//close := r.ReplaceAllString(match, "$4")
			attributes = strings.TrimSpace(attributes)

			if content == "" {
				if len(attributes) > 0 {
					return fmt.Sprintf("<%s %s/>", tag, attributes)
				}

				return fmt.Sprintf("<%s/>", tag)
			}

			return match
		},
	)
}
