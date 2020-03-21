package ab

import (
	"fmt"
	"testing"
)

func TestGetFT(t *testing.T) {
	t.Cleanup(func() {
		fmt.Println("cleaning up")
	})
}
