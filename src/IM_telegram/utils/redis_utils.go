package utils

import (
	"context"
	"fmt"
)

const (
	PublishKey = "websocket"
)

func Publish(ctx context.Context, channel string, mesg string) error {
	var err error
	fmt.Println("publish .....", mesg)
	err = RedisForIM.Publish(ctx, channel, mesg).Err()
	return err
}

func Subscribe(ctx context.Context, channel string) (string, error) {
	sub := RedisForIM.Subscribe(ctx, channel)
	fmt.Println("subscribe .....ctx ", ctx)
	msg, err := sub.ReceiveMessage(ctx)
	fmt.Println("subscribe .....msg.payload ", msg.Payload)
	return msg.Payload, err
}
