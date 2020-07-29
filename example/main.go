package main

import (
	gtjson "github.com/markd666/gtjson"
)

func main() {

	//gtjson.defaultIPAddress = "127.0.0.1"
	// create a gtjson client interface
	clientInterface := gtjson.CoreClient()
	clientInterface.Connect()

	if clientInterface.IsConnected() == true {
		telemetry := gtjson.GTTelemetry{
			PositionMeters: [3]float64{1.1, 2.2, 3.3},
			Quaterion:      [4]float64{0, 0, 0, 1},
			Timestamp:      12345,
		}
		clientInterface.SendTmToCore(telemetry)
	}
}
