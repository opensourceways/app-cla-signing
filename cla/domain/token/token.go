package token

type Token interface {
	New(interface{}) (string, error)
	Parse(string) (interface{}, error)
}
