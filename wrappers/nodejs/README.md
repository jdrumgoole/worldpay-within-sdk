# Node.JS Wrapper - Worldpay Within

This wrapper library allows Node.JS wrappers to natively interact with the Worldpay Within SDK Core via Thrift RPC. The thrift client implementation is abstracted from the developer and exposes a simple interface.

## Setup

* Required modules { `util`, `child_process`, `thrift` }
* Developed and test on Node.JS version 5.5.0
* It is is best to follow the sample application source code to get started with development.

## Running the test apps

In the `examples` directory there is a sample `sample-consumer` and `sample-producer` demonstrating how to use the SDK in a Producer and Consumer context.

Before starting the `main.js` Node.JS app file you must configure the RPC Agent. The main.js contains a line: similar to `client = wpwithin.createClient("127.0.0.1", 9090, function(err, response){`. While you can change the port value of 9090, the RPC Agent must be started using the same port, which is set in `conf.json`. Once the port has been set correctly you must start the RPC Agent via `./rpc-agent -configfile=conf.json`. You will need to leave this running while the Node app runs.

Once the RPC agent is running, you can now start the Node app with `node main.js`

** Please note the bundled RPC Agent is a Mac OS build. For alternate platforms please the `rpc-agent-bins` directory in the release folder **

### Next Steps
* Integration tests
* Create Node.JS package for [npm](https://www.npmjs.com/)

## The flows and API

[The flows and API can be found here](http://wptechinnovation.github.io/worldpay-within-sdk/the-flows.html)
