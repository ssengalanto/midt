package midt_test

import (
	"context"
	"testing"

	"github.com/ssengalanto/midt"
)

func Benchmark_Send(b *testing.B) {
	m := midt.New()

	hdlr := &CommandHandler{}
	err := m.RegisterRequestHandler(hdlr)
	if err != nil {
		b.Error(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := m.Send(context.Background(), &CommandRequest{})
		if err != nil {
			b.Error(err)
		}
	}
}

func Benchmark_Publish(b *testing.B) {
	m := midt.New()

	hdlr := &NotificationHandler{}
	err := m.RegisterNotificationHandler(hdlr)
	if err != nil {
		b.Error(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := m.Send(context.Background(), &NotificationRequest{})
		if err != nil {
			b.Error(err)
		}
	}
}
