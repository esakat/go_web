package backup

type Archiver interface {
	Archive(src, dist string) error
}
