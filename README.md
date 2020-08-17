# Ground Truth JSON converter
[![Go Report Card](https://goreportcard.com/badge/github.com/markd666/gtjson)](https://goreportcard.com/report/github.com/markd666/gtjson)
[![GoDoc](https://godoc.org/github.com/markd666/gtjson?status.svg)](https://godoc.org/github.com/markd666/gtjson)
[![Build Status](https://travis-ci.org/markd666/gtjson.svg?branch=master)](https://travis-ci.org/markd666/gtjson)

Used to convert GPS ground truth data into JSON format and send to core via TCP.

## Usage

`go get github.com/markd666/gtjson`


## Example

In the folder `/example` is a main.go file which shows the basic usuage of the package. 

`cd example`
`go run main.go`

or

`cd example`
`go build`
`./example.exe`

## Data Order

GTTelemetry struct has four variables. Note the order of x/y/z, 

```golang
PositionMeters [4]float64  // [0] = x | [1] = y | [2] = z
Quaterion      [4]float64  // [0] = w | [1] = x | [2] = y | [3] = z
Euler          [3]float64  // [0] = yaw | [1] = pitch | [2] = 'roll'
Timestamp      int64          // Unix Epoch time in milliseconds 
```