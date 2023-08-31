package huawei

import (
	"bytes"
	"encoding/json"

	"github.com/cnjack/cloudfunc/internal/provider/function"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	functiongraph "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/functiongraph/v2"
	functiongraphmodel "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/functiongraph/v2/model"
	functiongraphregion "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/functiongraph/v2/region"
)

var _ function.IProvider = (*Huawei)(nil)

func New(region string, credentials *basic.Credentials) *Huawei {
	client := functiongraph.NewFunctionGraphClient(
		functiongraph.FunctionGraphClientBuilder().
			WithRegion(functiongraphregion.ValueOf(region)).
			WithCredential(credentials).
			Build())
	return &Huawei{client}
}

type Huawei struct {
	*functiongraph.FunctionGraphClient
}

func (p *Huawei) Invoke(async bool, function_id string, req map[string]interface{}) (resp map[string]interface{}, request_id string, err error) {
	if async {
		hwResp, err := p.AsyncInvokeFunction(&functiongraphmodel.AsyncInvokeFunctionRequest{
			FunctionUrn: function_id,
			Body:        req,
		})
		if err != nil {
			return nil, "", err
		}
		return map[string]interface{}{}, *hwResp.RequestId, nil
	}
	hwResp, err := p.InvokeFunction(&functiongraphmodel.InvokeFunctionRequest{
		FunctionUrn: function_id,
		Body:        req,
	})
	if err != nil {
		return nil, "", err
	}
	err = json.NewDecoder(bytes.NewBufferString(*hwResp.Result)).Decode(&resp)
	if err != nil {
		return nil, "", err
	}
	return resp, *hwResp.RequestId, nil
}
