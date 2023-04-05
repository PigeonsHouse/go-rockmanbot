package reactions

import "strings"

func IsCalled(content string) bool {
	return strings.Contains(content, "ロックマン")
}
