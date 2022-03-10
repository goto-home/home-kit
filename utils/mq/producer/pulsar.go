package producer

import (
	"context"
	"github.com/apache/pulsar-client-go/pulsar"
)

type PulsarProducer struct {
	cli pulsar.Client
	pro pulsar.Producer
}

func (p *PulsarProducer) SendMsg(ctx context.Context, msg []byte) error {
	panic("implement me")
}

func (p *PulsarProducer) DelaySendMsg(ctx context.Context, msg []byte) error {
	panic("implement me")
}

func (p *PulsarProducer) CloseFunc() {
	p.pro.Close()
	p.cli.Close()
}
