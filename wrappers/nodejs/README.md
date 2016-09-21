# Node.js Wrapper - Worldpay Within

This wrapper library allows Node.js developers to natively interact with the Worldpay Within SDK Core via Thrift RPC. The thrift client implementation is abstracted from the developer and exposes a simple, familiar Node.js interface.

## Setup

* **Prerequisite**: Correctly installed and configured environment
* Required modules { `util`, `child_process`, `thrift` }
* Developed and tested on Node.js version 5.5.0
* It is is best to follow the sample application source code to get started with development.
* Ensure that the RPC-Agent binaries are downloaded from the release you are working with and put in `$WPWBIN` or `rpc-agent` directory in the working directory of the sample app.

## Running the example apps

In the `examples` directory there is `sample-consumer` and `sample-producer` demonstrating how to use the SDK in a Producer and Consumer context.

To run an example, navigate to the appropriate directory and use `node main.js`

**Note**: When running an example you will need to have both a producer and consumer running. Please ensure that both these applications are configured to talk to the RPC-Agent on different ports. Achieving this, is straightforward, please inspect the source of each main.js file

### Next Steps
* Integration tests
* Create Node.js package for [npm](https://www.npmjs.com/)

## The flows and API

[The flows and API can be found here](http://wptechinnovation.github.io/worldpay-within-sdk/the-flows.html)
