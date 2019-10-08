package cgomultiversion

import (
	"testing"
	cgomultiversion1 "github.com/vitalyisaev2/cgo_multiversion_go_lib/v1"
	cgomultiversion2 "github.com/vitalyisaev2/cgo_multiversion_go_lib/v2"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	// в первой версии библиотеки обычное сложение
	assert.Equal(t, 10, cgomultiversion1.Add(5, 5))
	// во второй версии библиотеки под сложениим скрывается умножение
	assert.Equal(t, 25, cgomultiversion2.Add(5, 5))
}
