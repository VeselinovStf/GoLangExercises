// Package reverse implements string revers functionalities
package reverse

// Reverse reverses a string
func Reverse(s string) string {
	sr := []rune(s)
	len := len(sr)
	res := make([]rune, len)
	for i := len - 1; i >= 0; i-- {
		res[len-1-i] = sr[i]
	}
	return string(res)
}
