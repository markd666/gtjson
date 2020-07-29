package gtjson

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
)

// GTTelemetry Ground Truth Telemetry data structure. This will be converted into JSON format and sent via
// TCP to the 'core'
type GTTelemetry struct {
	PositionMeters [3]float64 // [0] = x | [1] = y | [2] = z
	Quaterion      [4]float64 // [0] = w | [1] = x | [2] = y | [3] = z
	Timestamp      int32      // Unix Epoch time in milliseconds
}

type clientConnectionData struct {
	ipAddress   string
	port        int
	isConnected bool
	conn        net.Conn
}

// CoreClientInterface is an interface to the telemetry stream
type CoreClientInterface interface {
	Connect() error
	SendTmToCore(telemety GTTelemetry)
	GetPortNumber() int
	GetIPAddress() string
	IsConnected() bool
}

var (
	defaultIPAddress = "127.0.0.1"
	defaultPort      = 8899
)

// CoreClient returns an interface of type CoreClient
func CoreClient() CoreClientInterface {
	return &clientConnectionData{defaultIPAddress, defaultPort, false, nil}
}

// Connect creates a TCP server and blocks untilwaits for a client to connect
func (core *clientConnectionData) Connect() error {
	serverAddress := fmt.Sprintf("%v:%v", core.ipAddress, core.port)
	log.Printf("Starting TCP Server with address: %v", serverAddress)
	l, err := net.Listen("tcp", serverAddress)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer l.Close()

	var acceptErr error
	core.conn, acceptErr = l.Accept()
	if acceptErr != nil {
		fmt.Println(acceptErr)
		return acceptErr
	}

	core.isConnected = true

	return nil
}

// SendTmToCore
func (core *clientConnectionData) SendTmToCore(telemetry GTTelemetry) {
	// Convert gt_telemetry struct to json
	log.Println(telemetry)
	response, err := json.Marshal(telemetry)
	if err != nil {
		log.Fatalf("Unable to Marshal Telemtry to JSON format")
	}

	// send data over tcp link
	if core.IsConnected() == true {
		log.Println(string(response))
		core.conn.Write(response)
	}
}

// GetPortNumber returns as an integer the port number that the TCP server will use
func (core *clientConnectionData) GetPortNumber() int {
	return core.port
}

// GetIPAddress returns as a string the IP address that the TCP server will use
func (core *clientConnectionData) GetIPAddress() string {
	return core.ipAddress
}

// IsConnected returns as a boolean the status of the TCP connection
func (core *clientConnectionData) IsConnected() bool {
	return core.isConnected
}
