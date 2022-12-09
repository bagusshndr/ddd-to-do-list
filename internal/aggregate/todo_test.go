package aggregate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTodo(t *testing.T) {
	newTodo, _ := NewTodo(1, "bagus@bagus.com", 1, "high")
	t.Run("new todo", func(t *testing.T) {
		assert.NotNil(t, newTodo)
	})
}
