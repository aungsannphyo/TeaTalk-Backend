package sqlloader

import (
	"embed"

	"github.com/aungsannphyo/ywartalk/internal/store"
)

type EmbedSQLLoader struct {
	fs embed.FS
}

func NewEmbedSQLLoader() *EmbedSQLLoader {
	return &EmbedSQLLoader{fs: store.SqlFiles}
}

func (l *EmbedSQLLoader) LoadQuery(path string) (string, error) {
	queryBytes, err := l.fs.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(queryBytes), nil
}
