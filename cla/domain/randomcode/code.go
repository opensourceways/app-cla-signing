package randomcode

type RandomCode interface {
	New() (string, error)
	Validate(string) error
}
