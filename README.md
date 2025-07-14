# gedis

Trying to create a similar to Redis program in golang.

## Roadmap

- [x] GET,SET,DELETE
- [x] Persistance (Append-Only File)
- [ ] Binary protocol
- [ ] TTL

## Flags

- `-replay` It will replay the persistance file on startup (default: false)
- `-host` Host to bind the server (default: localhost)
- `-port` Port to bind the server (default: 8080)