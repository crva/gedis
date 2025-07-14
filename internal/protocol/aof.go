package protocol

import (
	"bufio"
	"os"

	"github.com/crva/gedis/internal/store"
)

type AOF struct {
	file *os.File
}

func NewAOF(filename string) (*AOF, error) {
	// Flags -> O_APPEND: Append to the end of the file
	// O_CREATE: Create the file if it does not exist
	// O_WRONLY: Open the file for writing only
	// Permissions -> 0644: Read and write for owner, read for group and others
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err // Return error if file cannot be opened or created
	}
	return &AOF{file: file}, nil // Return a new AOF instance with the opened file
}

func (aof *AOF) AppendGedisCommand(command string) error {
	_, err := aof.file.WriteString(command + "\n")
	if err != nil {
		return err
	}
	return nil
}

func (aof *AOF) Close() error {
	if aof.file != nil {
		return aof.file.Close()
	}
	return nil
}

func ReplayAOF(filename string, store *store.GedisStore) error {
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // If the file does not exist, return nil (no commands to replay)
		}
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		HandleCommand(line, store, nil)
	}
	return scanner.Err()
}
