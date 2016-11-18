package mqtt

// NOTE: Eclipse have revised their mqtt client API (finally!), but the following requires a rewrite to support it
/*
import (
	"log"
	"strconv"
)



import (
	"log"
	// Mqtt "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
	"strconv"
	"fmt"
)

type MqttPublisher struct {
	client *Mqtt.Client
}

var logIncoming Mqtt.MessageHandler = func(client *Mqtt.Client, msg Mqtt.Message) {
	log.Printf("MQTT [%s]: %s", msg.Topic(),msg.Payload())
}

func NewPublisher(addr, clientId string) *MqttPublisher {
	opts := Mqtt.NewClientOptions().AddBroker(addr)
	opts.SetClientID(client_id)
	opts.SetDefaultPublishHandler(logIncoming)
	pub := &MqttPublisher{
		client: Mqtt.NewClient(opts),
	}
	return pub
}

func (pub *MqttPublisher) Open() error {
	log.Printf("mqtt publisher connecting to serveer")
	if token := pub.client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("%v", token.Error())
		return token.Error()
	}
	log.Println("mqtt publisher is now connected")
	return nil
}

func (pub *MqttPublisher) Close() error {
	pub.client.Disconnect(250)
	return nil
}

func (pub *MqttPublisher) Publish(topic string,value interface{}) error {
	var msg string
	switch value := value.(type) {
	case string:
		msg = value
	case int:
		msg = strconv.Itoa(value)
	case float64:
		msg = strconv.FormatFloat(value, 'E', -1, 32)
	}

	if msg == "" {
		return fmt.Errorf("could not convert '%v'",value)
	}
	token := pub.client.Publish(topic, 0, false, msg)

	//token.Wait()
	// log.Printf("token result: ", token.Error())
	_ = token
	return nil
}

func (pub *MqttPublisher) PublishInt(topic string, i int) error {
	pub.client.Publish(topic, 0, false, strconv.Itoa(i))
	return nil
}

func (pub *MqttPublisher) PublishString(topic,value string) error {
	pub.client.Publish(topic, 0, false, value)
	return nil
}

func (pub *MqttPublisher) Flush() error {
	// noop
	return nil
}




*/