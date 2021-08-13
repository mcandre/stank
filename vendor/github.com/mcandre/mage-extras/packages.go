package mageextras

import (
	"strings"
)

// AllPackagesPath denotes all Go packages in a project.
var AllPackagesPath = strings.Join([]string{".", "..."}, PathSeparatorString)

// AllCommandsPath denotes all Go application packages in this project.
var AllCommandsPath = strings.Join([]string{".", "cmd", "..."}, PathSeparatorString)
