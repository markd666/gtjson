package gtjson

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func TestChangeIPAddress(t *testing.T) {
	device := CoreClient()

	device.SetIPAddress("192.168.0.1")

	if device.GetIPAddress() != "192.168.0.1" {
		t.Fatalf("Failed to set IP address to custom value")
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
			Timestamp:      1596103296,
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
		recvBuf := make([]byte, 0, 1024)
		tmp := make([]byte, 256)
		n, err := c.Read(tmp)
		if err != nil {
			break
		}

		recvBuf = append(recvBuf, tmp[:n]...)

		dataType := binary.BigEndian.Uint32(recvBuf[:4])
		if dataType != 13 {
			t.Fatalf("Incorrect data size")
		}

		dataSize := binary.BigEndian.Uint32(recvBuf[4:8])
		if dataSize != 91 {
			t.Fatalf("Incorrect message type")
		}

		var message GTTelemetry
		rawMessageBody := recvBuf[8:]
		error := json.Unmarshal(rawMessageBody, &message)
		if error != nil {

		}
		content, err := ioutil.ReadFile("data/" + "tm" + ".golden")
		if err != nil {
			t.Fatalf("Error loading golden file: %s", err)
		}
		want := string(content)
		got := string(rawMessageBody)
		if got != want {
			t.Errorf("Want:\n%s\nGot:\n%s", want, got)
		}

		return
	}
}
