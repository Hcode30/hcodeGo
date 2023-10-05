package hstrings

import "fmt"

func concat(str1, str2, sep string) string {

	return fmt.Sprintf(str1, sep, str2)
}
