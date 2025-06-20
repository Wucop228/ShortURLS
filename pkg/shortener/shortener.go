package shortener

const base = 62
const chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Generate(n int) string {
	if n == 0 {
		return "0"
	}
	res := ""
	for n > 0 {
		r := n % base
		res = string(chars[r]) + res
		n /= base
	}
	return res
}
