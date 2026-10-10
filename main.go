package main

import "fmt"

// These variables follow the Terraform provider release linker's conventional
// version/commit/date names. The current securefix fixture linker sets version.
var (
	version = "dev"
	commit  = "unknown"
	date    = "unknown"
)

func main() {
	fmt.Printf("test-only release fixture %s (%s, %s)\n", version, commit, date)
}
