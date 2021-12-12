package main

import "fmt"

var attempts = []string{
	"319",
	"680",
	"180",
	"690",
	"129",
	"620",
	"762",
	"689",
	"762",
	"318",
	"368",
	"710",
	"720",
	"710",
	"629",
	"168",
	"160",
	"689",
	"716",
	"731",
	"736",
	"729",
	"316",
	"729",
	"729",
	"710",
	"769",
	"290",
	"719",
	"680",
	"318",
	"389",
	"162",
	"289",
	"162",
	"718",
	"729",
	"319",
	"790",
	"680",
	"890",
	"362",
	"319",
	"760",
	"316",
	"729",
	"380",
	"319",
	"728",
	"716",
}

func hasSubstring(n int, a string) bool {
	s := fmt.Sprintf("%d", n)
	i := 0
	for p := 0; p < len(s) && i < len(a); p++ {
		if s[p] == a[i] {
			i++
		}
	}
	return i == len(a)
}

func main() {
	n := 100
	for {
		ok := true
		for i := 0; i < len(attempts); i++ {
			if !hasSubstring(n, attempts[i]) {
				ok = false
				break
			}
		}
		if ok {
			break
		}
		n++
	}
	fmt.Println(n)
}
