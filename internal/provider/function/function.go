package function

type IProvider interface {
	Invoke(async bool, functionID string, req map[string]interface{}) (resp map[string]interface{}, requestID string, err error)
}
