// raindrops package has a Convert function to convert a number to a string, the contents of which depend on the number's factors
package raindrops

import "fmt"

// Convert transform a number to a string, the contents of which depend on the number's factors
func Convert(n int) string {
	m := map[int]string{
		3: "Pling",
		5: "Plang",
		7: "Plong",
	}
	var r string
	for k, v := range m {
		if n%k == 0 {
			r += v
		}
	}

	if len(r) == 0 {
		return fmt.Sprintf("%v", n)
	}

	return r
}
