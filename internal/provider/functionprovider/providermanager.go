package functionprovider

import (
	"github.com/cnjack/cloudfunc/internal/provider/function"
)

type IProviderManager interface {
	Provider(name string) function.IProvider
}

type FunctionProviderManager map[string]function.IProvider

func MustGetFunctionProviderManager() IProviderManager {
	m := make(FunctionProviderManager)

	registeHuawei(m)

	return m
}

func (m FunctionProviderManager) Provider(name string) function.IProvider {
	return m[name]
}
