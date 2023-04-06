package email

import (
	"fmt"
	"testing"
)

func Test_Eamil(t *testing.T) {
	e, _ := Email()
	fmt.Printf("%s\n", e)
}
