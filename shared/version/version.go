package version

import "fmt"

var (
	Version    string
	CommitHash string
	BuiltAt    string
)

func GetVersion() string {
	return fmt.Sprintf("version %s commitHash=%s builtAt=%s",
		Version, CommitHash, BuiltAt)
}
