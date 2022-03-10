package mq

import "context"

type Producer interface {
	SendMsg(ctx context.Context, msg []byte) error
	DelaySendMsg(ctx context.Context, msg []byte) error
	CloseFunc()
}

type Consumer interface {
	Consume(ctx context.Context) ([]byte, error)
	// todo 延迟发送
}
