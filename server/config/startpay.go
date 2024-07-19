package config

type StartpayWeb3 struct {
	ApiKey      string `mapstructure:"api_key" json:"api_key" yaml:"api_key"`
	ApiSecret   string `mapstructure:"api_secret" json:"api_secret" yaml:"api_secret"`
	Host        string `mapstructure:"host" json:"host" yaml:"host"`
	ProjectPath string `mapstructure:"project_path" json:"project_path" yaml:"project_path"`
	SecretPath  string `mapstructure:"secre_path" json:"secre_path" yaml:"secre_path"`
}
