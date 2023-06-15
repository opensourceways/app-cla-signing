package domain

var config Config

func Init(cfg *Config) {
	config = *cfg
}

type Config struct {
	VerificationCodeExpiry  int64 `json:"verification_code_expiry"`
	MaxNumOfEmployeeManager int   `json:"max_num_of_employee_manager"`
}

func (cfg *Config) SetDefault() {
	if cfg.VerificationCodeExpiry <= 0 {
		cfg.VerificationCodeExpiry = 300
	}

	if cfg.MaxNumOfEmployeeManager <= 0 {
		cfg.MaxNumOfEmployeeManager = 5
	}
}
