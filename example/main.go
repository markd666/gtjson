package main

import (
	gtjson "github.com/markd666/gtjson"
)

func main() {

	// create a gtjson client interface
	clientInterface := gtjson.CoreClient()
	//Change the default IP Address from 127.0.0.1 to 192.168.0.102
	//clientInterface.SetIPAddress("192.168.0.102")
	clientInterface.Connect()

	if clientInterface.IsConnected() == true {
		telemetry := gtjson.GTTelemetry{
			PositionMeters: [3]float64{1.1, 2.2, 3.3},
			Quaterion:      [4]float64{0, 0, 0, 1},
			Euler:          [3]float64{10, 15, 20},
			Timestamp:      12345,
		}
		clientInterface.SendTmToCore(telemetry)
	}
}
