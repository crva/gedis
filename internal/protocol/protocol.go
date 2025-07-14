package protocol

import (
	"strings"

	"github.com/crva/gedis/internal/store"
)

func HandleCommand(command string, store *store.GedisStore) string {
	parts := strings.Split(command, ":")

	switch parts[0] {
	case "PING":
		return "PONG"
	default:
		return "ERROR: Unknown command"
	}
}
