package sdkmanCli

import (
	"os"
	"strings"
)

func isPathContains(candidate string) bool {
	return strings.Contains(os.Getenv("PATH"), candidate)
}
