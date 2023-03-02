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

	t.Run("new todo without title", func(t *testing.T) {
		todo, err := NewTodo(1, "", 1, "high")
		assert.NotNil(t, todo)
		assert.Error(t, err)
	})

	t.Run("new todos without title", func(t *testing.T) {
		todo, err := RebuildTodos(1, 1, "", 1, "high")
		assert.NotNil(t, todo)
		assert.Error(t, err)
	})

	t.Run("new todos without title", func(t *testing.T) {
		todo, _ := RebuildTodos(1, 1, "kerja bro", 1, "high")
		assert.NotNil(t, todo)
	})
}
