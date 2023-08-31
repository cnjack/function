package huawei

import "github.com/ory/viper"

type Config struct {
	AccessKey string
	SecretKey string
	Region    string
	Enable    bool
}

func NewConfig() *Config {
	v := viper.New()
	v.AutomaticEnv()
	v.SetEnvPrefix("HUAWEI")
	v.AddConfigPath(".")
	v.SetConfigFile("huawei.yaml")

	v.SetDefault("ENABLE", false)
	v.SetDefault("REGION", "cn-east-3")

	return &Config{
		AccessKey: v.GetString("ACCESS_KEY"),
		SecretKey: v.GetString("SECRET_KEY"),
		Region:    v.GetString("REGION"),
		Enable:    v.GetBool("ENABLE"),
	}
}
