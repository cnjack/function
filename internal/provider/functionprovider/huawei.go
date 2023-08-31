package functionprovider

import (
	huaweiauthorization "github.com/cnjack/cloudfunc/internal/provider/authorization/huawei"
	huaweifunctionprovider "github.com/cnjack/cloudfunc/internal/provider/function/huawei"
)

func registeHuawei(manager FunctionProviderManager) {
	config := huaweiauthorization.NewConfig()
	if !config.Enable {
		return
	}
	credentials := huaweiauthorization.NewHuaweiAuthorization(config)
	manager["huawei"] = huaweifunctionprovider.New(config.Region, credentials)
}
