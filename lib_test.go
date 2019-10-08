package cgomultiversion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	// в первой версии библиотеки обычное сложение
	assert.Equal(t, 10, AddV1(5, 5))
	// во второй версии библиотеки под сложениим скрывается умножение
	assert.Equal(t, 25, AddV2(5, 5))
}
