package repositoryimpl

type Config struct {
	Collections Collections `json:"collections" required:"true"`
}

type Collections struct {
	VerificationCode string `json:"verification_code" required:"true"`
}
