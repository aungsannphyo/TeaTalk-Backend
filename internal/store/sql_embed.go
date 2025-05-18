package store

import "embed"

//go:embed sql/**/*.sql
var SqlFiles embed.FS
