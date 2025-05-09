package main
import "fmt"
func main() {
	var n, x, hasil, c int
	var b bool
	b = false
	fmt.Scan(&n)
	visited := make(map[int]bool)
	for n > 0 && !b && !visited[n] {
		visited[n] = true
		c = n
		hasil = 0
		for c > 0 {
			x = c % 10
			x = x * x
			hasil += x
			c = c / 10
		}
		if hasil == 1 {
			b = true
		}
		n = hasil
	}
	fmt.Print(b)
}
