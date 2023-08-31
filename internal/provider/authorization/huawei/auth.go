package huawei

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
)

func NewHuaweiAuthorization(cfg *Config) *basic.Credentials {
	if !cfg.Enable {
		return nil
	}
	return basic.NewCredentialsBuilder().
		WithAk(cfg.AccessKey).
		WithSk(cfg.SecretKey).
		Build()
}
