package version

import "fmt"

var version string
var buildTime string

func Version() {
	fmt.Printf("Version: %s, Build Time: %s\n", version, buildTime)
	fmt.Println("Author: feeeei unipay")
}
