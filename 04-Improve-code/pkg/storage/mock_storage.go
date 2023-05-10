package storage

type MockStorage struct {
	SpyGetvalue bool
	Data        map[string]interface{}
}

func (mk *MockStorage) GetValue(key string) interface{} {
	mk.SpyGetvalue = true
	if v, ok := mk.Data[key]; ok {
		return v
	}
	return nil
}

func NewMockStorage() *MockStorage {
	return &MockStorage{}
}
