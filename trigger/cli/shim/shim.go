/*
Sniperkit-Bot
- Status: analyzed
*/

package main

import (
	"fmt"
	"os"

	"github.com/sniperkit/snk.fork.tibcosoftware-flogo-contrib/trigger/cli"
)

func main() {
	result, _ := cli.Invoke()
	fmt.Fprintf(os.Stdout, "%s", result)
	os.Exit(0)
}
