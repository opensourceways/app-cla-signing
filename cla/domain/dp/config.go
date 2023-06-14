package dp

var config Config

func Init(cfg *Config) {
	config = *cfg
}

type Config struct {
	MaxLengthOfCorpName int `json:"max_length_of_corp_name"   required:"true"`
}
