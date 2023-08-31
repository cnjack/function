package huawei

import "github.com/spf13/viper"

type Config struct {
	AccessKey string
	SecretKey string
	Region    string
	Enable    bool
}

func NewConfig() *Config {
	v := viper.New()
	v.AutomaticEnv()
	v.AddConfigPath(".")
	v.SetConfigFile("config.yaml")

	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}

	v.SetDefault("HUAWEI.ENABLE", false)
	v.SetDefault("HUAWEI.REGION", "cn-east-3")

	return &Config{
		AccessKey: v.GetString("HUAWEI.ACCESS_KEY"),
		SecretKey: v.GetString("HUAWEI.SECRET_KEY"),
		Region:    v.GetString("HUAWEI.REGION"),
		Enable:    v.GetBool("HUAWEI.ENABLE"),
	}
}
