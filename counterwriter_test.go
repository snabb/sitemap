package sitemap_test

import (
	"fmt"
	"os"

	"github.com/snabb/sitemap"
)

func ExampleCounterWriter() {
	cw := sitemap.NewCounterWriter(os.Stdout)
	fmt.Fprintln(cw, "hello world")
	fmt.Println(cw.Count())
	// Output:
	// hello world
	// 12
}
