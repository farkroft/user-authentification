package mock

type MockConfig struct{}

func (c *MockConfig) GetString(str string) string {
	return ""
}
