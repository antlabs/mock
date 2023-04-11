package country

import (
	"fmt"
	"testing"
)

func TestCountry(t *testing.T) {
	t.Parallel()
	// t.Skip("skip test")
	for i := 0; i < 100; i++ {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			i := i
			t.Parallel()
			// t.Skip("skip test")
			name := Country(i%2 == 0)
			if name == "" {
				t.Fatal("name is empty")
			}
		})
	}

}

func TestCountry2(t *testing.T) {
	name := Country(true)
	fmt.Printf("%s\n", name)
	name = Country(false)
	fmt.Printf("%s\n", name)
}
