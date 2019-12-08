package config

type Config struct {
	BAAK struct {
		Host                 string
		JadkulEndpoint       string
		RetryMaxAttempt      int
		RetryIntervalSeconds int
	} `gcfg:"BAAK"`
}
