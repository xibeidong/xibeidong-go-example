package Template

import "testing"

func TestNewWorker(t *testing.T) {
	w := NewWorker(&Coder{})
	w.Daily()
}
