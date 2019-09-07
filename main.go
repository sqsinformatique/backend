// main
package main

import (
	"fmt"

	"github.com/sqsinformatique/backend/cmd"
	"github.com/sqsinformatique/backend/utils"
)

const defAppFieldValue = "undefined"

var (
	name    = defAppFieldValue
	version = defAppFieldValue
	builded = defAppFieldValue
	hash    = defAppFieldValue

	appInfo = fmt.Sprintf(
		"%s. version: %s, builded: %s, hash: %s",
		name, version, builded, hash,
	)
)

func main() {
	utils.Info(appInfo)

	cmd.Execute()
}
