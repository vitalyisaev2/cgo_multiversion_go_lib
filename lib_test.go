package cgomultiversion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	assert.Equal(t, 4, Add(2, 2))
}
