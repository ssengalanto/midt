package mediatr_test

import (
	"github.com/stretchr/testify/require"
	"testing"

	mediatr "github.com/ssengalanto/mediatR"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run("it should create a new mediatr instance", func(t *testing.T) {
		m := mediatr.New()
		assert.NotNil(t, m, "new mediatr instance shouldn't be nil")
	})
}

func TestMediatr_RegisterRequestHandler(t *testing.T) {
	t.Run("it should register all handlers successfully", func(t *testing.T) {
		m := mediatr.New()

		hdlr1 := &CommandHandler{}
		err := m.RegisterRequestHandler(hdlr1)
		require.NoError(t, err)

		hdlr2 := &QueryHandler{}
		err = m.RegisterRequestHandler(hdlr2)
		require.NoError(t, err)
	})
	t.Run("it should return an error when trying to register an already existing handler", func(t *testing.T) {
		m := mediatr.New()

		hdlr1 := &CommandHandler{}
		err := m.RegisterRequestHandler(hdlr1)
		require.NoError(t, err)

		hdlr2 := &CommandHandler{}
		err = m.RegisterRequestHandler(hdlr2)
		require.Error(t, err)
	})
}
