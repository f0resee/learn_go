package rocketmq

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/admin"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"testing"
)

var topicName = "testTopic"
var endPoint = []string{"127.0.0.1:10909"}

func TestRMQ(t *testing.T) {
	t.Run("rocketmq create topic", func(t *testing.T) {
		topic := "newTwo"
		//clusterName := "DefaultCluster"
		nameSrvAddr := []string{"127.0.0.1:9876"}
		brokerAddr := "127.0.0.1:10911"

		testAdmin, err := admin.NewAdmin(admin.WithResolver(primitive.NewPassthroughResolver(nameSrvAddr)))
		if err != nil {
			fmt.Println(err.Error())
		}

		//create topic
		err = testAdmin.CreateTopic(
			context.Background(),
			admin.WithTopicCreate(topic),
			admin.WithBrokerAddrCreate(brokerAddr),
		)
		if err != nil {
			fmt.Println("Create topic error:", err.Error())
		}

		//deletetopic
		err = testAdmin.DeleteTopic(
			context.Background(),
			admin.WithTopicDelete(topic),
			//admin.WithBrokerAddrDelete(brokerAddr),
			//admin.WithNameSrvAddr(nameSrvAddr),
		)
		if err != nil {
			fmt.Println("Delete topic error:", err.Error())
		}

		err = testAdmin.Close()
		if err != nil {
			fmt.Printf("Shutdown admin error: %s", err.Error())
		}
	})

	t.Run("send message", func(t *testing.T) {
		p, _ := rocketmq.NewProducer(
			producer.WithNameServer(endPoint),
			producer.WithRetry(2),
			producer.WithGroupName("ProducerGroupName"),
		)

		err := p.Start()
		if err != nil {
			t.Log("start producer error: ", err.Error())
		}

		message := "hello"
		result, err := p.SendSync(context.Background(), &primitive.Message{
			Topic: topicName,
			Body:  []byte(message),
		})
		if err != nil {
			t.Log(err)
		} else {
			t.Log(result)
		}
	})

}
