package domain

var config Config

func Init(cfg *Config) {
	config = *cfg
}

type Config struct {
	VerificationCodeExpiry int64 `json:"verification_code_expiry"`
}

func (cfg *Config) SetDefault() {
	if cfg.VerificationCodeExpiry == 0 {
		cfg.VerificationCodeExpiry = 300
	}
}
