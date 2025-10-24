package hamming
import "fmt"
func Distance(a, b string) (int, error) {
   if len(a) != len(b) {
		return 0, fmt.Errorf("sample must be of equal length, got %d and %d", len(a), 	len(b))
	}
	if len(a) == 0 {
		return 0, nil
	}

	diff := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			diff++
		}
	}

	return diff, nil
}
