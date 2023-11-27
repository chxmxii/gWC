package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDisplayLogo(t *testing.T) {
	assert.NotPanics(t, func() {
		err := DisplayLogo(WithLogoPath("../../hack/ascii.txt"))
		assert.Nil(t, err)
	})
}
