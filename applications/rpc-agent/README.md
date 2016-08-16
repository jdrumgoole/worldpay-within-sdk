# RPC-AGENT
An application that hosts the WorldpayWithin SDK behind an Apache Thrift RPC listener. This agent and the SDK are written in Go but Thrift allows for other languages to interact with the SDK via it's RPC mechanism.

## Usage

Start the application and specify a port using the `-port` flag. e.g. `rpc-agent -port=9090` or `.rpc-agent -port 9090`

## Notes

* The file currently outputs at a debug level to `output.log`.
* Currently working on full configurability of this application. Configuration will be possible via both the command line and also a configuration file.
