package function

type IProvider interface {
	Invoke(async bool, function_id string, req map[string]interface{}) (resp map[string]interface{}, request_id string, err error)
}
