# worldpay-within-sdk
Worldpay Within SDK to allow payments within IoT.

The core of this SDK is written in Go with a native Go interface. Along with the native Go interface is an RPC layer (Apache Thrift) to allow communication through other languages. It is intended that we will develop a number of complementary wrapper libraries for other languages which should include C#.NET, Java, Python, Node.JS at a minimum.

### Top level directories

* applications - Applications used to support this SDK.
* rpc - Thrift definitions for the RPC layer.
* sdkcore - Worldpay Within SDK Core written in go.
* wrappers - Wrapper implementations in other languages using Thrift RPC.

## Further documentation

* [Please see our documentation pages for more details on what Worldpay Within is](http://wptechInnovation.github.io/worldpay-within-sdk)
* [Also for a detailed architecture guide for Worldpay Within please see our full documentation](http://wptechinnovation.github.io/worldpay-within-sdk/architecture.html)

## The SDK binaries - if you don't don't want to compile from sources

### Binary builds

Please see the releases section of GitHub for access to pre-built binaries of the RPC Agent and Dev Client apps.

Both of the apps have been built for 32bit and 64bit architectures on Windows, MacOS, Linux and Linux (ARM).

While the RPC Agent can be run from anywhere it makes sense to add it to the directory of the application which will call it. The reason for this is that each application needs it own running instance of the RPC Agent and it can easily get confusing when there are multiple applications and agents deployed.

Please see the examples in both the Java and Node.JS wrappers. these examples currently hold binaries for MacOS x64. I suggest replacing that binary with an alternate build, if required.

The dev client can really be run from anywhere as it is not coupled with anything else.

### Usage

* RPC Agent `./rpc-agent -configfile <filename>`. Please see explanation of rpc-agent config file for further info.

* Dev Client `./dev-client`

## How to use this SDK

If you intend to develop a Go application then you need not concern yourself with the RPC interface or any wrapper libraries, these are only required if you wish to work in another language.

To develop using Go you must use the package `wpwithin` in the `sdkcore` directory.

If you wish to develop using a wrapper library then please navigate to your chosen language from the `wrappers` directory. Please see the included sample code on how to consume the SDK.

### Go development

* Install Go command line tool
* This is pretty simple to find on windows and Mac OS (OS X), note that Linux can be slightly trickier, but shouldn't be a problem either
* Set up the environmental variables correctly; you only need to set $GOPATH, and that should be set as `<home>/<required_path>/<cloned_repo_structure>`, where `<home>` is wherever you want the code, `<required_path>` is `/src/github.com/wptechinnovation/`. For clarification the `$GOPATH` variable is where the source code for the Worldpay Within SDK will be, or you applicaiton code will reside, and is not the location of the actual golang binaries (for compiling and running go).
* clone the repo to `$GOPATH/src/github.com/wptechinnovation/`
* Get the dependencies; `go get ./...`

### Install the RPC agent
* Change directory to `cd $GOPATH/src/github.com/wptechninnovation/worldpay-within-sdk/applications/rpc-agent`
* Type `go install`
* This should build, package up, and install the binaries for the rpc-agent into your bin directory `$GOPATH/bin`
* If there are any errors around missing packages do additional `go get <package-repo-path>`
* If there are any compile errors, it is likely you are running a version of go that is too old (we have seen this most commonly on Ubuntu Linux)

### Install the example client app
* Change directory to `$GOPATH/src/github.com/wptechninnovation/worldpay-within-sdk/applications/dev-client/`
* Type `go install`
* This should build, package up, and install the binaries for the rpc-agent into your bin directory `$GOPATH/bin`
* If there are any errors around missing packages do additional `go get <package-repo-path>`
* If there are any compile errors, it is likely you are running a version of go that is too old (we have seen this most commonly on Ubuntu Linux)

### Run the RPC agent
* Running the RPC agent is critical to the SDK, if working with GOLANG, then it will need to be manually run or kicked off automatically by your app
* Running any of the wrappers, then the RPC agent will be automatically started for you, however this may not be available in early releases, and so you should be aware of how to manually run the RPC agent yourself
* Change to the bin directory `cd $GOPATH/bin`
* Type the following command to run the RPC agent and see the command line flags that you can pass; `./rpc-agent -help`
* You can manually set the parameters, to get everything running quickly you just need to set the prot e.g. `./rpc-agent -port 9090`
* Alternatively you can use the configuration file provided to configure the RPC agent, to do this type; `./rpc-agent -configfile <path and filename of config file>`
* We have provided a config file in the source directory so; `./rpc-agent -configfile $GOPATH/src/worldpay-within-sdk/applications/rpc-agent/conf.json`

### Run the Client app
* Change to the bin directory `cd $GOPATH/bin`
* Type the following command to run the client app; `./dev-client`
* The dev client should start, and you should be able to interact with the menu, more details on operating the dev-client app can be found under the dev-client app folder or the full documentation pages on for Worldpay Within on github.
