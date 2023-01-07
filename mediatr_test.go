package mediatr_test

import (
	"testing"

	mediatr "github.com/ssengalanto/mediatR"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run("create new mediatr instance", func(t *testing.T) {
		m := mediatr.New()
		assert.NotNil(t, m, "new mediatr instance shouldn't be nil")
	})
}
