package sqlloader

type LoaderType int

const (
	EmbedLoader LoaderType = iota
)

func SQLLoaderFactory(loaderType LoaderType) SQLLoader {
	switch loaderType {
	case EmbedLoader:
		return NewEmbedSQLLoader()
	default:
		panic("unsupported loader type")
	}
}
