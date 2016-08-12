# worldpay-within-sdk
Worldpay Within SDK to allow payments within IoT.

The core of this SDK is written in Go with a native Go interface. Along with the native Go interface is an RPC layer (Apache Thrift) to allow communication through other languages. It is intended that we will develop a number of complementary wrapper libraries for other languages which should include C#.NET, Java, Python at a minimum.

### Install

* Install Go command line
* Set up the environmental variables correctly; you only need to set $GOPATH, and that should be set as <home>/<required_path>/<cloned_repo_structure>, where <home> is wherever you want the code, <required_path> is /src/innovation.worldpay.com
* clone the repo to $GOPATH/src/innovation.worldpay.com
* Get the dependencies; go get github.com/Sirupsen/logrus
* Get the dependencies; go get github.com/gorilla/mux
* Get the dependencies; go get github.com/nu7hatch/gouuid
* Get the dependencies; go get git.apache.org/thrift.git/lib/go/thrift


### Configuration file versus command line flags
The RPC client takes command line flags e.g. -port 9091 but it can also take the flag -configfile 'conf.json' so you can specify the configuration in a config file. For example

```
{
        "WorldpayWithinConfig": {
                "BufferSize" : 100,
                "Buffered": false,
                "Framed": false,
                "Host": "127.0.0.1",
                "Logfile": "worldpayWithin.log",
                "Loglevel": "warn",
                "Port": 9081,
                "Protocol": "binary",
                "Secure": false
        }
}
```


# Initial pre-alpha release - June 6, 2016


* Core SDK somewhat complete but not 100%. No service handover (begin/end)
* Thrift definition of SDK service and message types
* Basic Java program demonstrating RPC function
* RPC Agent tool to enable starting the RPC from command line and programmatically. All options exposed via CLI flags. use -h for usage.
* C# namespace TBD (A.Brodie)
* BUG: There is an issue with the int->price map in the Thrift services definition (Pointer/Value error in Go). This has been disabled for now.
* Only binary transport works in Java/Go RPC example. Will investigate others.
* Added a semi implemented console application (dev-client) which shows the usage of the SDK in Go. This is probably the best documentation for now :)


# Next steps
