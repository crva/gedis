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
	case "SET":
		if len(parts) != 3 {
			return "ERROR: SET command requires a key and a value"
		}
		store.Set(parts[1], parts[2])
		return "OK"
	case "GET":
		if len(parts) != 2 {
			return "ERROR: GET command requires a key"
		}
		value, exists := store.Get(parts[1])
		if !exists {
			return "ERROR: Key not found"
		}
		return value
	case "DEL":
		if len(parts) != 2 {
			return "ERROR: DEL command requires a key"
		}
		store.Delete(parts[1])
		return "OK"
	case "KEYS":
		keys := store.Keys() // Assuming a Keys method exists in GedisStore
		return strings.Join(keys, ";")
	default:
		return "ERROR: Unknown command"
	}
}
