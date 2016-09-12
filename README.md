# worldpay-within-sdk
Worldpay Within SDK to allow payments within IoT.

The core of this SDK is written in Go with a native Go interface. Along with the native Go interface is an RPC layer (Apache Thrift) to allow communication through other languages.

Currently, there are wrappers available for the following technologies:
* Node.js
* .NET
* Java
* Python (in development)

**Note 1**: Please note that if you intend to work with one of the wrapper frameworks, it is not required that you build the Go source code directly. With each release we will bundle pre-built binaries of the RPC-Agent application. The RPC-Agent is an application that starts the Thrift RPC interface into the Go SDK Core. Once this application is up and running the wrapper can communicate with the SDK Core. In the latest release of the SDK, the RPC-Agent is started automatically by the wrapper.

**Note 2**: To enable payments for your instance of the SDK and applications, you will need to create an account at [Worldpay Online Payments](online.worldpay.com). Once the account is created, please navigate to *settings* -> *API Keys* and keep note of the *service key* and *client key* for later. You will need to add these keys into your sample apps when "intialising a producer".

### Top level directories

* applications - Applications used to support this SDK.
* rpc - Thrift definitions for the RPC layer.
* sdkcore - Worldpay Within SDK Core written in Go.
* wrappers - Wrapper implementations in other languages using Thrift RPC.

## Further documentation

* [Please see our documentation pages for more details on what Worldpay Within is](http://wptechInnovation.github.io/worldpay-within-sdk)
* [Also for a detailed architecture guide for Worldpay Within please see our full documentation](http://wptechinnovation.github.io/worldpay-within-sdk/architecture.html)

## The SDK binaries - if you don't don't want to compile from sources

### Binary builds

Please see the releases section of GitHub for access to pre-built binaries of the RPC Agent and Dev Client apps.

Both of the apps have been built for 32bit and 64bit architectures on Windows, MacOS, Linux and Linux ARM.

To enable the example wrapper applications, please use put the prebuilt binaries in a folder `rpc-agent` at the root level of the sample application. Alternative, you can put the binaries in a directory that the environment variable `WPWBIN` points to.

### Example Usage

* RPC Agent `./rpc-agent -port 9099 -logfile=wpwithin.log -loglevel=debug,warn,info,error,fatal -callbackport=9098`

* Dev Client `./dev-client`

## How to use this SDK

If you intend to develop a Go application then you need not concern yourself with the RPC interface or any wrapper libraries, these are only required if you wish to work in another language.

To develop using Go you must use the package `wpwithin` in the `sdkcore` directory. Please see the `examples` directory in `sdkcore`.

If you wish to develop using a wrapper library then please navigate to your chosen language from the `wrappers` directory and see the included sample source code and readme files.

### Go development
* Prerequisite: correctly installed and configured environment
* `go get github.com/wptechinnovation/worldpay-within-sdk` will download the SDK to your $GOPATH
* Install Go dependencies: `cd applications/rpc-agent` then run: `go get ./...`

### Install the RPC agent
* Change directory to `cd $GOPATH/src/github.com/wptechninnovation/worldpay-within-sdk/applications/rpc-agent`
* Type `go install`
* This should build, package up, and install the binary for the rpc-agent into your bin directory `$GOPATH/bin`
* If there are any errors around missing packages do additional `go get <package-repo-path>`
* If there are any compile errors, it is likely you are running a version of go that is too old (we have seen this most commonly on Ubuntu Linux)

### Install the example client app
* Change directory to `$GOPATH/src/github.com/wptechninnovation/worldpay-within-sdk/applications/dev-client/`
* Type `go install`
* This should build, package up, and install the binary for the rpc-agent into your bin directory `$GOPATH/bin`
* If there are any errors around missing packages do additional `go get <package-repo-path>`
* If there are any compile errors, it is likely you are running a version of go that is too old (we have seen this most commonly on Ubuntu Linux)

### Run the RPC agent
* Running the RPC agent is critical to the SDK, if working with GOLANG, then it will need to be manually run or kicked off automatically by your app
* Running any of the wrappers, then the RPC agent will be automatically started for you, however this may not be available in early releases, and so you should be aware of how to manually run the RPC agent yourself
* Change to the bin directory `cd $GOPATH/bin`
* Type the following command to run the RPC agent and see the command line flags that you can pass; `./rpc-agent -help`
* You can manually set the parameters, to get everything running quickly you just need to set the port e.g. `./rpc-agent -port 9090`
* Alternatively you can use the configuration file provided to configure the RPC agent, to do this type; `./rpc-agent -configfile <path and filename of config file>`

### Run the Client app
* Change to the bin directory `cd $GOPATH/bin`
* Type the following command to run the client app; `./dev-client`
* The dev client should start, and you should be able to interact with the menu, more details on operating the dev-client app can be found under the dev-client app folder or the full documentation pages on for Worldpay Within on github.
