package config

// Config defines the configuration structure.
type Config struct {
	General struct {
		LogLevel               int  `mapstructure:"log_level"`
		LogToSyslog            bool `mapstructure:"log_to_syslog"`
		PasswordHashIterations int  `mapstructure:"password_hash_iterations"`
	} `mapstructure:"general"`

	Gateway struct {
		Url     string `mapstructure:"url"`
		CACert  string `mapstructure:"ca_cert"`
		TLSCert string `mapstructure:"tls_cert"`
		TLSKey  string `mapstructure:"tls_key"`
	} `mapstructure:"gateway"`

	Rest struct {
		Bind string `mapstructure:"bind"`
		Url  string `mapstructure:"url"`
	} `mapstructure:"rest"`
}

// C holds the global configuration.
var C Config
