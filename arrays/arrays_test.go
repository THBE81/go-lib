package arrays

import "testing"

func TestGetFirstOfType(t *testing.T) {
	t.Run("Get first integer element", func(t *testing.T) {
		arr := Array("A string", 69, "Your mom", 420, nil, any(0))
		result, _ := GetFirstOfType[any, int](arr)
		expected := 69
		if result != expected {
			t.Errorf("TestGetFirstOfType()=%d, want %d", result, expected)
		}
	})
}
