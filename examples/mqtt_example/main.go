package main

import (
	"log"
	"github.com/fuzzycow/ev32go/ev3api"
	"github.com/fuzzycow/ev32go/ev3api/device/sensor"
	"github.com/fuzzycow/ev32go/clip"
	"github.com/fuzzycow/ev32go/helpers/monitor"
	"time"
	Mqtt "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
)

var addr string = "tcp://192.168.1.10:1883"
var mqtt_client *Mqtt.Client

func prepareIR() *sensor.Infrared {
	ir := clip.NewInfraredSensor(ev3api.INPUT_1)
	if err := ir.Open(); err != nil {
		log.Fatalf("Failed to open device: %v", err)
	}
	return ir
}
func prepareMqttClient() *Mqtt.Client {
	opts := Mqtt.NewClientOptions().AddBroker(addr)
	opts.SetClientID("ev32go")
	opts.SetDefaultPublishHandler(logIncoming)

	log.Printf("Connecting to %s",addr)
	//create and start a client using the above ClientOptions
	c := Mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("%v",token.Error())
	}

	log.Println("connected")

	return c
	//subscribe to the topic /go-mqtt/sample and request messages to be delivered
	//at a maximum qos of zero, wait for the receipt to confirm the subscription
	/*
	if token := c.Subscribe("go-mqtt/sample", 0, nil); token.Wait() && token.Error() != nil {
		log.Fatalf("%v",token.Error())
	}
	*/

	//Publish 5 messages to /go-mqtt/sample at qos 1 and wait for the receipt
	//from the server after sending each message

	/*
	//unsubscribe from /go-mqtt/sample
	if token := c.Unsubscribe("go-mqtt/sample"); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
	*/

	//c.Disconnect(250)
}

var logIncoming Mqtt.MessageHandler = func(client *Mqtt.Client, msg Mqtt.Message) {
	log.Printf("TOPIC: %s\n", msg.Topic())
	log.Printf("MSG: %s\n", msg.Payload())
}



func main() {
	c := prepareMqttClient()
	ir := clip.NewInfraredSensor(ev3api.INPUT_1)
	if err := ir.Open(); err != nil {
		log.Fatalf("Failed to open device: %v", err)
	}
	ir.SetMode(ir.Mode_IR_PROX())

	mon := monitor.New(monitor.Everything, 50 * time.Millisecond, 1000 * time.Second )
	defer mon.Stop()

	distCh := mon.PollString(func() string {
		return  ir.GetAttrString("value0")
	})

	log.Printf("main loop running")
	MONITORING: for {
		select {
		case dist, ok := <-distCh:
			if !ok {
				break MONITORING
			}
			//log.Printf("publishing...")
			token := c.Publish("/sensor/ir", 0, false, dist)
			//log.Printf("waiting for ack...")
			token.Wait()
			//log.Printf("got ack.")
		}
	}
}

