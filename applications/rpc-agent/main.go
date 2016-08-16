package main

import (
	"flag"
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin"
	conf "github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin/configLoad"
	"github.com/wptechinnovation/worldpay-within-sdk/sdkcore/wpwithin/rpc"
)

/*
	A simple program to enable the WPWithin Core RPC interface.
	The intention is that this program is called by language wrappers in order to gain RPC access to the core.
*/

// Application exit codes
const EXIT_OK = 0
const EXIT_GENERAL_ERR = 1

// Log level selectors passed in as argument
const LEVEL_PANIC = "panic"
const LEVEL_FATAL = "fatal"
const LEVEL_ERROR = "error"
const LEVEL_WARN = "warn"
const LEVEL_INFO = "info"
const LEVEL_DEBUG = "debug"

// General constants
const LOGFILE_PERMS = 0755
const RPC_MIN_PORT = 1
const DEFAULT_ARG_CONFIGFILE = ""
const DEFAULT_ARG_PORT = 0 // Defaulting to this should cause error (desired) (force port specifer// )
const DEFAULT_ARG_TRANSPORT_BUFFER = 8192
const DEFAULT_ARG_FRAMED = false
const DEFAULT_ARG_BUFFERED = false
const DEFAULT_ARG_SECURE = false
const DEFAULT_ARG_HOST = "127.0.0.1"
const DEFAULT_ARG_PROTOCOL = "binary"

const ARG_NAME_CONFIGFILE = "configfile"
const ARG_NAME_LOG_LEVEL = "loglevel"
const ARG_NAME_LOGFILE = "logfile"
const ARG_NAME_PROTOCOL = "protocol"
const ARG_NAME_FRAMED = "framed"
const ARG_NAME_BUFFERED = "buffered"
const ARG_NAME_HOST = "host"
const ARG_NAME_PORT = "port"
const ARG_NAME_SECURE = "secure"
const ARG_NAME_BUFFER = "buffer"

// Globally scoped vars
var sdk wpwithin.WPWithin
var rpcConfig rpc.Configuration

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
	os.Exit(EXIT_OK)
}

func initArgs() {

	log.Debug("Begin initArgs()")

	// Determine whether to use config file
	configFilePtr := flag.String(ARG_NAME_CONFIGFILE, DEFAULT_ARG_CONFIGFILE, "Config file name - string.")

	// Log config args
	logLevelPtr := flag.String(ARG_NAME_LOG_LEVEL, LEVEL_WARN, "Log level")
	logFilePtr := flag.String(ARG_NAME_LOGFILE, "", "Log file, if set, outputs to file, if not, not logfile.")

	// Program specific arguments
	protocolPtr := flag.String(ARG_NAME_PROTOCOL, DEFAULT_ARG_PROTOCOL, "Transport protocol.")
	framedPtr := flag.Bool(ARG_NAME_FRAMED, DEFAULT_ARG_FRAMED, "Framed transmission - bool.")
	bufferedPtr := flag.Bool(ARG_NAME_BUFFERED, DEFAULT_ARG_BUFFERED, "Buffered transmission - bool.")
	hostPtr := flag.String(ARG_NAME_HOST, DEFAULT_ARG_HOST, "Listening host.")
	portPtr := flag.Int(ARG_NAME_PORT, DEFAULT_ARG_PORT, "Port to listen on. Required.")
	securePtr := flag.Bool(ARG_NAME_SECURE, DEFAULT_ARG_SECURE, "Secured transport - bool.")
	bufferPtr := flag.Int(ARG_NAME_BUFFER, DEFAULT_ARG_TRANSPORT_BUFFER, "Buffer size.")

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

	log.Debug("After flag.parse()")

	logLevelValue := *logLevelPtr
	logFileValue := *logFilePtr

	if "" != configFileValue {
		log.Debug("Begin PopulateConfiguration() from config file")

		// Pull from config file - command line overwrites
		rpcConfig = rpc.Configuration{}
		rpcConfig = conf.PopulateConfiguration(configFileValue, rpcConfig)

		log.Debug("End PopulateConfiguration() from config file")

		// Use config file
		logLevelValue = rpcConfig.Loglevel
		logFileValue = rpcConfig.Logfile

		// Program specific arguments
		protocolValue = rpcConfig.Protocol
		framedValue = rpcConfig.Framed
		bufferedValue = rpcConfig.Buffered
		hostValue = rpcConfig.Host
		portValue = rpcConfig.Port
		secureValue = rpcConfig.Secure
		bufferValue = rpcConfig.BufferSize

		log.Debug("Before parsing the config file")
		// TODO write parser for config file
		log.Debug("After parsing the config file")

	}

	log.Debug("Before log setup")

	log.Debug("Begin parsing log level arguments")

	switch logLevelValue {

	case LEVEL_PANIC:
		log.SetLevel(log.PanicLevel)
	case LEVEL_FATAL:
		log.SetLevel(log.FatalLevel)
	case LEVEL_ERROR:
		log.SetLevel(log.ErrorLevel)
	default:
	case LEVEL_WARN:
		log.SetLevel(log.WarnLevel)
	case LEVEL_INFO:
		log.SetLevel(log.InfoLevel)
	case LEVEL_DEBUG:
		log.SetLevel(log.DebugLevel)
	}

	log.Debug("Begin parsing log level arguments")

	log.Debug("Begin parsing log file arguments and setup log file")
	if logFileValue != "" {

		log.WithField("File", logFileValue).Debug("Will logs to file.")

		logFile, err := os.OpenFile(logFileValue, os.O_WRONLY|os.O_CREATE, LOGFILE_PERMS)

		if err != nil {

			log.Warn(fmt.Sprintf("Could not create log file", err.Error()))
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
	log.Debug("After assign RPC config.")

	log.Debug("End initArgs()")
}

func startRPC() {

	log.Debug("Before startRPC()")

	// Validate required (with no defaults)
	if rpcConfig.Port < RPC_MIN_PORT {

		log.WithFields(log.Fields{"Port": rpcConfig}).Fatal("Invalid listening port provided")

		fmt.Println("Port value must be greater than zero")

		os.Exit(EXIT_GENERAL_ERR)
	}

	log.WithField("Configuration: ", fmt.Sprintf("%+v", rpcConfig)).Debug("Before rpc.NewService")
	rpc, err := rpc.NewService(rpcConfig, sdk)
	log.Debug("After rpc.NewService")

	if err != nil {

		fmt.Printf("Error create new RPC service: %q\n", err.Error())

		log.WithFields(log.Fields{"Error": err.Error()}).Fatal("Error creating new RPC service")

		os.Exit(EXIT_GENERAL_ERR)
	}

	log.WithFields(log.Fields{"port": rpcConfig.Port}).Debug("Attempting to start RPC interface on port")
	if err := rpc.Start(); err != nil {

		log.WithFields(log.Fields{"Error": err.Error()}).Fatal("Error starting RPC service")

		fmt.Printf("Error starting RPC service: %q\n", err.Error())

		os.Exit(EXIT_GENERAL_ERR)
	}

	log.Debug("End startRPC()")
}
