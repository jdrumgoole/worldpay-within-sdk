package main

import (
	"flag"
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin"
	"github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin/configuration"
	"github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin/rpc"
)

/*
	A simple program to enable the WPWithin Core RPC interface.
	The intention is that this program is called by language wrappers in order to gain RPC access to the core.
*/

const exitOK = 0
const exitGeneralErr = 1

// Log level selectors passed in as argument
const levelPanic = "panic"
const levelFatal = "fatal"
const levelError = "error"
const levelWarn = "warn"
const levelInfo = "info"
const levelDebug = "debug"

// General constants
const logfilePerms = 0755
const rpcMinPort = 1
const defaultArgConfigFile = ""
const defaultArgPort = 0 // Defaulting to this should cause error (desired) (force port specifer// )
const defaultArgTransportBuffer = 8192
const defaultArgFramed = false
const defaultArgBuffered = false
const defaultArgSecure = false
const defaultArgHost = "127.0.0.1"
const defaultArgProtocol = "binary"
const defaultArgCallbackPort = 0 // Default 0 means callback feature not to be used

const argNameConfigFile = "configfile"
const argNameLogLevel = "loglevel"
const argNameLogfile = "logfile"
const argNameProtocol = "protocol"
const argNameFramed = "framed"
const argNameBuffered = "buffered"
const argNameHost = "host"
const argNamePort = "port"
const argNameSecure = "secure"
const argNameBuffer = "buffer"
const argNameCallbackPort = "callbackport"

// Globally scoped vars
var sdk wpwithin.WPWithin
var rpcConfig rpc.Configuration

const (
	keyBufferSize string = "BufferSize"
	keyBuffered   string = "Buffered"
	keyFramed     string = "Framed"
	keyHost       string = "Host"
	keyLogfile    string = "Logfile"
	keyLoglevel   string = "Loglevel"
	keyPort       string = "Port"
	keyProtocol   string = "Protocol"
	keySecure     string = "Secure"
	keyCallbackPort string = "CallbackPort"
)

func main() {

	// Start off by setting logging to a high level
	// This way we can catch output during initial setup of args and logging via arguments.
	log.SetLevel(log.DebugLevel)

	log.Debug("Begin main()")

	log.Debug("Before initArgs().")
	initArgs()
	log.Debug("After initArgs().")

	log.Debug("Before startRPC()")
	startRPC()
	log.Debug("After startRPC()")

	fmt.Println("Program end...")
	log.Debug("Application end - End main()")

	os.Exit(exitOK)
}

func initArgs() {

	log.Debug("Begin initArgs()")

	// Determine whether to use config file
	configFilePtr := flag.String(argNameConfigFile, defaultArgConfigFile, "Config file name - string.")

	// Log config args
	logLevelPtr := flag.String(argNameLogLevel, levelWarn, "Log level")
	logFilePtr := flag.String(argNameLogfile, "", "Log file, if set, outputs to file, if not, not logfile.")

	// Program specific arguments
	protocolPtr := flag.String(argNameProtocol, defaultArgProtocol, "Transport protocol.")
	framedPtr := flag.Bool(argNameFramed, defaultArgFramed, "Framed transmission - bool.")
	bufferedPtr := flag.Bool(argNameBuffered, defaultArgBuffered, "Buffered transmission - bool.")
	hostPtr := flag.String(argNameHost, defaultArgHost, "Listening host.")
	portPtr := flag.Int(argNamePort, defaultArgPort, "Port to listen on. Required.")
	securePtr := flag.Bool(argNameSecure, defaultArgSecure, "Secured transport - bool.")
	bufferPtr := flag.Int(argNameBuffer, defaultArgTransportBuffer, "Buffer size.")
	callbackPortPtr := flag.Int(argNameCallbackPort, defaultArgCallbackPort, "Callback Port")

	log.Debug("Before flag.parse()")
	flag.Parse()

	configFileValue := *configFilePtr
	protocolValue := *protocolPtr
	framedValue := *framedPtr
	bufferedValue := *bufferedPtr
	hostValue := *hostPtr
	portValue := *portPtr
	secureValue := *securePtr
	bufferValue := *bufferPtr
	callbackPortValue := *callbackPortPtr

	log.Debug("After flag.parse()")

	logLevelValue := *logLevelPtr
	logFileValue := *logFilePtr

	if "" != configFileValue {
		log.Debug("Begin PopulateConfiguration() from config file")

		// Pull from config file - command line overwrites
		// rpcConfig := conf.PopulateConfiguration(configFileValue)
		config, err := configuration.Load(configFileValue)

		if err != nil {

			fmt.Println(err.Error())
			os.Exit(2)
		}

		log.Debug("End PopulateConfiguration() from config file")

		// Use config file
		logLevelValue = config.GetValue(keyLoglevel).Value
		logFileValue = config.GetValue(keyLogfile).Value

		// Program specific arguments
		protocolValue = config.GetValue(keyProtocol).Value
		framed, err := config.GetValue(keyFramed).ReadBool()
		framedValue = framed
		buffered, err := config.GetValue(keyBuffered).ReadBool()
		bufferedValue = buffered
		hostValue = config.GetValue(keyHost).Value
		port, err := config.GetValue(keyPort).ReadInt()
		portValue = port
		secure, err := config.GetValue(keySecure).ReadBool()
		secureValue = secure
		bufferSize, err := config.GetValue(keyBufferSize).ReadInt()
		bufferValue = bufferSize
		callbackPort, err := config.GetValue(keyCallbackPort).ReadInt()
		callbackPortValue = callbackPort

		log.Debug("Before parsing the config file")
		// TODO write parser for config file
		log.Debug("After parsing the config file")

	}

	log.Debug("Before log setup")

	log.Debug("Begin parsing log level arguments")

	switch logLevelValue {

	case levelPanic:
		log.SetLevel(log.PanicLevel)
	case levelFatal:
		log.SetLevel(log.FatalLevel)
	case levelError:
		log.SetLevel(log.ErrorLevel)
	default:
	case levelWarn:
		log.SetLevel(log.WarnLevel)
	case levelInfo:
		log.SetLevel(log.InfoLevel)
	case levelDebug:
		log.SetLevel(log.DebugLevel)
	}

	log.Debug("Begin parsing log level arguments")

	log.Debug("Begin parsing log file arguments and setup log file")
	if logFileValue != "" {

		log.WithField("File", logFileValue).Debug("Will logs to file.")

		logFile, err := os.OpenFile(logFileValue, os.O_WRONLY|os.O_CREATE, logfilePerms)

		if err != nil {

			log.Warn(fmt.Sprintf("Could not create log file: %s", err.Error()))
		} else {

			log.Debug("Setting up log text formatter")
			tf := &log.TextFormatter{}
			tf.DisableColors = true
			tf.FullTimestamp = true
			log.SetFormatter(tf)
			log.WithField("TextFormatter", fmt.Sprintf("%+v", tf)).Debug("End set up log text formatter")

			log.Debug("Successfully created log file.. setting output now.")
			log.SetOutput(logFile)
		}

	} else {

		log.Debug("Will not be logging to file - no logfile argument detected.")
	}
	log.Debug("End parsing log file arguments and setup log file")

	log.Debug("After log setup")

	log.Debug("Before assign RPC config.")
	rpcConfig.Protocol = protocolValue
	rpcConfig.Framed = framedValue
	rpcConfig.Buffered = bufferedValue
	rpcConfig.Host = hostValue
	rpcConfig.Port = portValue
	rpcConfig.Secure = secureValue
	rpcConfig.BufferSize = bufferValue
	rpcConfig.CallbackPort = callbackPortValue
	log.Debug("After assign RPC config.")

	log.Debug("End initArgs()")
}

func startRPC() {

	log.Debug("Before startRPC()")

	// Validate required (with no defaults)
	if rpcConfig.Port < rpcMinPort {

		log.WithFields(log.Fields{"Port": rpcConfig}).Fatal("Invalid listening port provided")

		fmt.Println("Port value must be greater than zero")

		os.Exit(exitGeneralErr)
	}

	log.WithField("Configuration: ", fmt.Sprintf("%+v", rpcConfig)).Debug("Before rpc.NewService")
	rpc, err := rpc.NewService(rpcConfig, sdk)
	log.Debug("After rpc.NewService")

	if err != nil {

		fmt.Printf("Error create new RPC service: %q\n", err.Error())

		log.WithFields(log.Fields{"Error": err.Error()}).Fatal("Error creating new RPC service")

		os.Exit(exitGeneralErr)
	}

	log.WithFields(log.Fields{"port": rpcConfig.Port}).Debug("Attempting to start RPC interface on port")
	if err := rpc.Start(); err != nil {

		log.WithFields(log.Fields{"Error": err.Error()}).Fatal("Error starting RPC service")

		fmt.Printf("Error starting RPC service: %q\n", err.Error())

		os.Exit(exitGeneralErr)
	}

	log.Debug("End startRPC()")
}
