package tokenimpl

func NewTokenImpl() *tokenImpl {
	return &tokenImpl{}
}

type tokenImpl struct{}

func (impl *tokenImpl) New(interface{}) (string, error) { return "", nil }

func (impl *tokenImpl) Parse(string) (interface{}, error) { return nil, nil }
