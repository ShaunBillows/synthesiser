package signal

type MockOutput struct {
    data []float64
}

func NewMockOutput() *MockOutput {
	return &MockOutput{}
}

func (m *MockOutput) Write(data []float64) {
    m.data = data
}
