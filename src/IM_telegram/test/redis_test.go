package test

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"sync"
	"testing"
	"time"
)

var wg sync.WaitGroup
var RedisForIM = redis.NewClient(&redis.Options{
	Addr:         "120.46.203.178:6379",
	Password:     "",
	DB:           0,
	PoolSize:     30,
	MinIdleConns: 30,
})

func TestRedisBaseOption(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	// 执行命令获取结果
	val, err := RedisForIM.Get(ctx, "key").Result()
	fmt.Println(val, err)

	// 先获取到命令对象
	cmder := RedisForIM.Get(ctx, "key")
	fmt.Println(cmder.Val()) // 获取值
	fmt.Println(cmder.Err()) // 获取错误

	// 直接执行命令获取错误
	err = RedisForIM.Set(ctx, "key", 10, time.Hour).Err()

	// 直接执行命令获取值
	value := RedisForIM.Get(ctx, "key").Val()
	fmt.Println(value)

	RedisForIM.Del(ctx, "key")
	err = RedisForIM.Get(ctx, "key").Err()
	t.Log(err)

	err = RedisForIM.Do(ctx, "set", "key", 10, "EX", 3600).Err()
	if err != nil {
		t.Log(err)
	}
	result, err := RedisForIM.Do(ctx, "get", "key").Result()
	if err != nil {
		t.Log(err)
	}
	t.Log(result)
}

func TestRedisOrderSetOption(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	for i := 1; i <= 10; i++ {
		err := RedisForIM.ZAdd(ctx, "orderSet", &redis.Z{
			Score:  float64(i),
			Member: fmt.Sprintf("its %d yo", i),
		}).Err()
		if err != nil {
			t.Log(err)
		}
	}

	result, err := RedisForIM.ZCard(ctx, "orderSet").Result()
	if err != nil {
		t.Log(err)
	}
	t.Log("ZCard :", result)

	strings, err := RedisForIM.ZRange(ctx, "orderSet", 0, -1).Result()
	if err != nil {
		t.Log(err)
	}
	for _, v := range strings {
		t.Log("ZRange :", v)
	}

	strings, err = RedisForIM.ZRevRange(ctx, "orderSet", 0, -1).Result()
	if err != nil {
		t.Log(err)
	}
	for _, v := range strings {
		t.Log("ZRevRange :", v)
	}
}

func subscribe(ctx context.Context, key string) {
	topic := RedisForIM.Subscribe(ctx, key)
	message, err := topic.ReceiveMessage(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(message.Channel, "       ", message.Payload)
	redisSubChannel := topic.Channel()
	for mesg := range redisSubChannel {
		fmt.Println(mesg.Channel, "       ", mesg.Payload)
	}
	wg.Done()
}

func TestMesgSubAndPub(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	go subscribe(ctx, "OneForAll")
	go subscribe(ctx, "OneForAll")
	go subscribe(ctx, "OneForAll")
	time.Sleep(time.Second)
	for i := 1; i <= 4; i++ {
		RedisForIM.Publish(ctx, "OneForAll", fmt.Sprintf("its %d test publish yo", i))
	}

	//go subscribe(ctx,"OneForAll")
	//go subscribe(ctx,"OneForAll")
	//go subscribe(ctx,"OneForAll")
	for {
		time.Sleep(time.Second)
	}
}
