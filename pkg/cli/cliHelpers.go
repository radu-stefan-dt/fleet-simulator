package cli

import (
	"strings"
)

func parseFlagEnvironment(env string) string {
	env = strings.TrimPrefix(env, "https://")
	env = strings.TrimSuffix(env, "/")
	return env
}
func parseFlagNumFleets(nf int) int {
	switch {
	case nf < 1:
		return 1
	case nf > 10:
		return 10
	default:
		return nf
	}
}
