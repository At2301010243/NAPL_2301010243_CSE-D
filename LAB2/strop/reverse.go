// reversefunction.go
package strop

func Reverse(s string) string {
	rev := ""
	for i := len(s) - 1; i >= 0; i-- {
		rev = rev + string(s[i])
	}
	return rev
}
