package producer

import (
	"context"
	"github.com/apache/pulsar-client-go/pulsar"
	"time"
)

type PulsarProducer struct {
	cli pulsar.Client
	pro pulsar.Producer
}

func NewPulsarProducer(url, topic string) *PulsarProducer {
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL: url,
	})
	if err != nil {
		panic(err)
	}

	pul := &PulsarProducer{cli: client}
	// producer
	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: topic,
	})
	if err != nil {
		panic(err)
	}
	pul.pro = producer
	return pul
}

func (p *PulsarProducer) SendMsg(ctx context.Context, msg []byte) error {
	_, err := p.pro.Send(ctx, &pulsar.ProducerMessage{
		Payload: msg,
	})
	return err
}

func (p *PulsarProducer) DelayAtSendMsg(ctx context.Context, msg []byte, date *time.Time) error {
	_, err := p.pro.Send(ctx, &pulsar.ProducerMessage{
		Payload:   msg,
		DeliverAt: *date,
	})
	return err
}

func (p *PulsarProducer) DelayAfterSendMsg(ctx context.Context, msg []byte, dur time.Duration) error {
	_, err := p.pro.Send(ctx, &pulsar.ProducerMessage{
		Payload:      msg,
		DeliverAfter: dur,
	})
	return err
}

func (p *PulsarProducer) CloseFunc() {
	p.pro.Close()
	p.cli.Close()
}
