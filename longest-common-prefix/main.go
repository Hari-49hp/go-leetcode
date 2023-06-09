package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
}

func longestCommonPrefix(strs []string) string {

	if 1 <= len(strs) && len(strs) <= 200 {

		prefix := strs[0]

		for i := 0; i < len(strs); i++ {

			for !strings.HasPrefix(strs[i], prefix) {
				prefix = prefix[:len(prefix)-1]
			}

		}

		return prefix
	}

	return ""
}
