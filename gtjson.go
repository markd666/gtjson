package gtjson

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
)

type gtTelemetry struct {
	positionMetres [3]float64
	quaterion      [4]float64
	timestamp      float64
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
}

const (
	defaultIPAddress = "192.168.0.100"
	defaultPort      = 8899
)

// CoreClient returns an interface of type CoreClient
func CoreClient() CoreClientInterface {
	return &clientConnectionData{defaultIPAddress, defaultPort, false, nil}
}

func (core *clientConnectionData) Connect() error {
	serverAddress := fmt.Sprintf("%v:%v", core.ipAddress, core.port)
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
	return nil
}

func (core *clientConnectionData) SentTmToCore(telemetry gtTelemetry) {
	// Convert gt_telemetry struct to json
	telmetryJSON, err := json.Marshal(telemetry)
	if err != nil {
		log.Fatalf("Unable to Marshal Telemtry to JSON format")
	}

	// send data over tcp link
}
