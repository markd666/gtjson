package gtjson

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"sync"
	"testing"
)

func TestConstructor(t *testing.T) {

	device := CoreClient()

	if device.GetPortNumber() != 8899 {
		t.Fatalf("Using incorrect port number")
	}

	if device.GetIPAddress() != "127.0.0.1" {
		t.Fatalf("Using incorrect IP address")
	}

}

func TestBasicClientConnection(t *testing.T) {
	go basicClientConnection()

	defaultIPAddress = "127.0.0.1"
	// create a gtjson client interface
	clientInterface := CoreClient()
	clientInterface.Connect()

	if clientInterface.IsConnected() != true {
		t.Fatalf("Expected client interface to indicate a connection had been established")
	}

}

func basicClientConnection() {
	connectAddress := "127.0.0.1:8899"
	c, err := net.Dial("tcp", connectAddress)
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Printf("Basic Client connected to: %v", c.RemoteAddr())
}

func TestServerWrite(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	go clientConnection(t, &wg)

	defaultIPAddress = "127.0.0.1"
	// create a gtjson client interface
	clientInterface := CoreClient()
	clientInterface.Connect()

	if clientInterface.IsConnected() == true {
		telemetry := GTTelemetry{
			PositionMeters: [3]float64{1.1, 2.2, 3.3},
			Quaterion:      [4]float64{0, 0, 0, 1},
			Timestamp:      12345,
		}
		clientInterface.SendTmToCore(telemetry)
	}

	wg.Wait()

}

func clientConnection(t *testing.T, wg *sync.WaitGroup) {
	defer wg.Done()

	connectAddress := "127.0.0.1:8899"
	c, err := net.Dial("tcp", connectAddress)
	if err != nil {
		fmt.Println(err)
		return
	}
	t.Logf("Client connected to: %v", c.RemoteAddr())

	for {

		d := json.NewDecoder(c)

		var message GTTelemetry
		err := d.Decode(&message)
		if err != nil {

		}
		t.Log(message)

		if message.PositionMeters[0] != float64(1.1) {
			t.Fatalf("PositionMeters[0] incorrect")
		}

		if message.PositionMeters[1] != float64(2.2) {
			t.Fatalf("PositionMeters[1] incorrect")
		}

		if message.PositionMeters[2] != float64(3.3) {
			t.Fatalf("PositionMeters[2] incorrect")
		}

		return

	}
}
