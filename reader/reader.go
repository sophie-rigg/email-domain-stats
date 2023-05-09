package reader

type Reader interface {
	Read() (map[string]int, error)
}
