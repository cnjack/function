package provider

import (
	huaweicredential "github.com/cnjack/cloudfunc/internal/provider/authorization/huawei"
	"github.com/cnjack/cloudfunc/internal/provider/function/huawei"
	"testing"
)

func Test_Huawei(t *testing.T) {
	config := huaweicredential.NewConfig()
	credential := huaweicredential.NewHuaweiAuthorization(config)
	huaweiFuncClient := huawei.New(config.Region, credential)
	resp, requestID, err := huaweiFuncClient.Invoke(false, "urn:fss:cn-north-4:251f06bc9cf5441f8f4b70e117583f8b:function:default:test:latest", map[string]interface{}{})
	t.Log(resp)
	t.Log(requestID)
	t.Log(err)
}
