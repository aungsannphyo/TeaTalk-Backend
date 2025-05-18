package sqlloader

type SQLLoader interface {
	LoadQuery(path string) (string, error)
}
