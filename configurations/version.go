package configurations

import _ "embed"

//go:embed version.txt
var version []byte

func Version() string {
	return string(version)
}