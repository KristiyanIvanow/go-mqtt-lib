# go-mqtt-lib

A comprehensive Go SDK for MQTT-based IoT Edge container communication. This library provides client implementations for direct method invocations, container properties management, and MQTT message handling.

## Installation

```bash
go get github.com/KristiyanIvanow/go-mqtt-lib@v1.0.0
```

## Features

- **MQTT Client** (`mqttclient`): Singleton MQTT broker client with automatic reconnection
- **Direct Method Handler** (`directmethod`): Handle direct method invocations via MQTT
- **Container Properties** (`containerproperties`): Manage reported and desired properties
- **Type System** (`types`): Enums, constants, and shared types
- **Models** (`models`): Data structures for configuration, messages, and state
- **Logger** (`logger`): Pluggable logging interface
- **Error Handling** (`errors`): Custom SDK error types

## Quick Start

### Initialize MQTT Client

```go
package main

import (
	"github.com/KristiyanIvanow/go-mqtt-lib/src/models"
	"github.com/KristiyanIvanow/go-mqtt-lib/src/mqttclient"
)

func main() {
	config := &models.MqttConfigModel{
		BrokerAddr: "tcp://localhost:1883",
		// ... other config
	}
	
	client := mqttclient.GetInstance()
	client.Init(config, nil)
	defer client.Close()
	
	// Client is ready to use
}
```

### Register Direct Method Handler

```go
import "github.com/KristiyanIvanow/go-mqtt-lib/src/directmethod"

dmClient := directmethod.GetInstance()
dmClient.RegisterMethodHandler("myMethod", func(payload interface{}) (directmethod.DirectMethodResponse, error) {
	return directmethod.DirectMethodResponse{
		Status:  200,
		Payload: "Success",
	}, nil
})
```

### Manage Container Properties

```go
import "github.com/KristiyanIvanow/go-mqtt-lib/src/containerproperties"

propsClient := containerproperties.GetInstance()
propsClient.ReportProperties(map[string]interface{}{
	"temperature": 25.5,
	"status":      "active",
})
```

## Project Structure

```
├── src/
│   ├── sdk.go                          # SDK root package
│   ├── types/                          # Enums and shared types
│   ├── models/                         # Data models
│   ├── errors/                         # Custom error types
│   ├── logger/                         # Logger interface
│   ├── mqttapi/                        # Message API interface
│   ├── mqttclient/                     # MQTT client implementation
│   ├── directmethod/                   # Direct method handler
│   └── containerproperties/            # Container properties manager
├── go.mod                              # Module definition
└── go.sum                              # Dependency checksums
```

## Importing Subpackages

To import individual packages from this library:

```go
// Types and constants
import "github.com/KristiyanIvanow/go-mqtt-lib/src/types"

// Models
import "github.com/KristiyanIvanow/go-mqtt-lib/src/models"

// MQTT Client
import "github.com/KristiyanIvanow/go-mqtt-lib/src/mqttclient"

// Direct Methods
import "github.com/KristiyanIvanow/go-mqtt-lib/src/directmethod"

// Container Properties
import "github.com/KristiyanIvanow/go-mqtt-lib/src/containerproperties"

// Error Types
import "github.com/KristiyanIvanow/go-mqtt-lib/src/errors"

// Logger
import "github.com/KristiyanIvanow/go-mqtt-lib/src/logger"
```

## Dependencies

- Go 1.18+
- `github.com/eclipse/paho.mqtt.golang` - MQTT client library

## License

See LICENSE file for details.

## Contributing

Contributions are welcome! Please ensure all code passes tests and follows Go conventions.

## Version

v1.0.0
