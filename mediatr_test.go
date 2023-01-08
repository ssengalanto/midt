package mediatr_test

import (
	"context"
	"testing"

	mediatr "github.com/ssengalanto/mediatR"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	t.Run("it should create a new mediatr instance", func(t *testing.T) {
		m := mediatr.New()
		assert.NotNil(t, m, "new mediatr instance shouldn't be nil")
	})
}

func TestMediatr_RegisterRequestHandler(t *testing.T) {
	t.Run("it should register all request handlers successfully", func(t *testing.T) {
		m := mediatr.New()

		hdlr1 := &CommandHandler{}
		err := m.RegisterRequestHandler(hdlr1)
		require.NoError(t, err)
	})
	t.Run("it should return an error when trying to register an already existing request handler", func(t *testing.T) {
		m := mediatr.New()

		hdlr1 := &CommandHandler{}
		err := m.RegisterRequestHandler(hdlr1)
		require.NoError(t, err)

		hdlr2 := &CommandHandler{}
		err = m.RegisterRequestHandler(hdlr2)
		require.Error(t, err)
	})
}

func TestMediatr_RegisterNotificationHandler(t *testing.T) {
	t.Run("it should register the notification handler successfully", func(t *testing.T) {
		m := mediatr.New()

		hdlr := &NotificationHandler{}
		err := m.RegisterNotificationHandler(hdlr)
		require.NoError(t, err)
	})
	t.Run("it should return an error when trying to register an already existing notification handler", func(t *testing.T) {
		m := mediatr.New()

		hdlr1 := &NotificationHandler{}
		err := m.RegisterNotificationHandler(hdlr1)
		require.NoError(t, err)

		hdlr2 := &NotificationHandler{}
		err = m.RegisterNotificationHandler(hdlr2)
		require.Error(t, err)
	})
}

func TestMediatr_RegisterPipelineBehaviour(t *testing.T) {
	t.Run("it should register the pipeline behaviour successfully", func(t *testing.T) {
		m := mediatr.New()

		pb := &PipelineBehaviourHandler{}
		err := m.RegisterPipelineBehaviour(pb)
		require.NoError(t, err)
	})
	t.Run("it should return an error when trying to register an already existing pipeline behaviour",
		func(t *testing.T) {
			m := mediatr.New()

			pb1 := &PipelineBehaviourHandler{}
			err := m.RegisterPipelineBehaviour(pb1)
			require.NoError(t, err)

			pb2 := &PipelineBehaviourHandler{}
			err = m.RegisterPipelineBehaviour(pb2)
			require.Error(t, err)
		})
}

func TestMediatr_Send(t *testing.T) {
	t.Run("it should execute the request handler",
		func(t *testing.T) {
			m := mediatr.New()

			hdlr := &CommandHandler{}
			err := m.RegisterRequestHandler(hdlr)
			require.NoError(t, err)

			_, err = m.Send(context.Background(), &CommandRequest{})
			require.NoError(t, err)
		})
	t.Run("it should execute the pipeline behaviours",
		func(t *testing.T) {
			m := mediatr.New()

			pb := &PipelineBehaviourHandler{}
			err := m.RegisterPipelineBehaviour(pb)
			require.NoError(t, err)

			hdlr := &CommandHandler{}
			err = m.RegisterRequestHandler(hdlr)
			require.NoError(t, err)

			_, err = m.Send(context.Background(), &CommandRequest{})
			require.NoError(t, err)
		})
	t.Run("it should return an error if there are no registered request handlers in the registry",
		func(t *testing.T) {
			m := mediatr.New()

			_, err := m.Send(context.Background(), &CommandRequest{})
			require.Error(t, err)
		})
}

func TestMediatr_Publish(t *testing.T) {
	t.Run("it should execute the notification handler",
		func(t *testing.T) {
			m := mediatr.New()

			hdlr := &NotificationHandler{}
			err := m.RegisterNotificationHandler(hdlr)
			require.NoError(t, err)

			err = m.Publish(context.Background(), &NotificationRequest{})
			require.NoError(t, err)
		})
	t.Run("it should return an error if there are no registered notification handlers in the registry",
		func(t *testing.T) {
			m := mediatr.New()

			err := m.Publish(context.Background(), &NotificationRequest{})
			require.Error(t, err)
		})
}
