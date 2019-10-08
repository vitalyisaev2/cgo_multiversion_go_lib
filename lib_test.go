package cgomultiversion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	// во второй версии библиотеки под сложениим скрывается умножение
	assert.Equal(t, 25, Add(5, 5))
}
