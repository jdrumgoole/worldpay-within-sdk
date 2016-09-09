#SDK Core package (GOLANG)

This is the actual source code for the core Worldpay Within SDK it is separated into the following paths;

* `wpwithin` - This contains the actual source code for the core code of Worldpay Within
* `wpwithin_test` - contains tests to various parts of the SDK

## WPWITHIN - core code

The core Worldpay Within code, does the important parts of the work, including host terminal emulation, contacts to the payment gateway, managing the RPC calls, and service discovery, etc.

### configuration

This enables the core to load in key information from a config file.

### core

This contains the core and the factory. The factory creates all the core objects required by the core. The core holds references to all the critical objects. Or to be more precise, the SDK Core acts as a container for dependencies of the SDK.

### hte

The client allows interaction with the HTE service, it contains the HTE service and the service handler which Coordinates requests between RPC interface and internal SDK interface. There is a credential store for the 'terminal' or payments gateway credentials. The Order Manager coordinates during negotitation, payment and delivery flows. There are also some help http request object(s).

* [More detail about the flows and associate diagrams can be found in the detailed documentation here](http://wptechinnovation.github.io/worldpay-within-sdk/architecture.html)

### psp

This code enables the Worldpay Within SDK to communicate with the online.worldpay.com payments gatesway to make payments. The results of these payment can be viewed by having an associated account (associated with the credentials) on online.worldpay.com.

### rpc

This manages the services and objects that are exposed by thrift.

### servicediscovery

This contains the broadcaster, scanner and communicator. The broadcaster and scanner allows the UDP messages to be broadcast (sent out) or scanned (received). The communicator allows comms over various layers.

### types

These are the go objects containing the important data that makes up the Worldpay Within world of objects that are required and that interact.

### utils

Various utilities for networking, text processing, UUID processing, etc.
