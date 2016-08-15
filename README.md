# worldpay-within-sdk
Worldpay Within SDK to allow payments within IoT.

The core of this SDK is written in Go with a native Go interface. Along with the native Go interface is an RPC layer (Apache Thrift) to allow communication through other languages. It is intended that we will develop a number of complementary wrapper libraries for other languages which should include C#.NET, Java, Python, Node.JS at a minimum.

### Top level directories

* applications - Applications used to support this SDK.
* rpc - Thrift definitions for the RPC layer.
* sdkcore - Worldpay Within SDK Core written in go.
* wrappers - Wrapper implementations in other languages using Thrift RPC.

## How to use this SDK

If you intend to develop a Go application then you need not concern yourself with the RPC interface or any wrapper libraries, these are only required if you wish to work in another language.

To develop using Go you must use the package `wpwithin` in the `sdkcore` directory.

If you wish to develop using a wrapper library then please navigate to your chosen language from the `wrappers` directory. Please see the included sample code on how to consume the SDK.

### Go development

* Install Go command line
* Set up the environmental variables correctly; you only need to set $GOPATH, and that should be set as <home>/<required_path>/<cloned_repo_structure>, where <home> is wherever you want the code, <required_path> is /src/innovation.worldpay.com
* clone the repo to $GOPATH/src/innovation.worldpay.com
* Get the dependencies; go get github.com/Sirupsen/logrus
* Get the dependencies; go get github.com/gorilla/mux
* Get the dependencies; go get github.com/nu7hatch/gouuid
* Get the dependencies; go get git.apache.org/thrift.git/lib/go/thrift
