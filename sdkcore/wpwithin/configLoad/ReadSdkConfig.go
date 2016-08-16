package configLoad

import (
	"encoding/json"
	"fmt"
	"os"

	rpc "innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/rpc"
)

type IndividualConfig struct {
	BufferSize int
	Buffered   bool
	Framed     bool
	Host       string
	Logfile    string
	Loglevel   string
	Port       int
	Protocol   string
	Secure     bool
}

type ConfigurationLocal struct {
	WorldpayWithinConfig IndividualConfig
}

func loadConfig(configPath string) (configuration ConfigurationLocal) {

	if configPath == "" {

		configPath = "conf.json"
	}

	file, _ := os.Open(configPath)
	decoder := json.NewDecoder(file)
	configuration = ConfigurationLocal{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	return
}

func PopulateConfiguration(configPath string, rpcConfig rpc.Configuration) (rpcConfigReturn rpc.Configuration) {

	configuration := loadConfig(configPath)

	rpcConfig.Protocol = configuration.WorldpayWithinConfig.Protocol
	rpcConfig.Framed = configuration.WorldpayWithinConfig.Framed
	rpcConfig.Buffered = configuration.WorldpayWithinConfig.Buffered
	rpcConfig.Logfile = configuration.WorldpayWithinConfig.Logfile
	rpcConfig.Host = configuration.WorldpayWithinConfig.Host
	rpcConfig.Port = configuration.WorldpayWithinConfig.Port
	rpcConfig.Secure = configuration.WorldpayWithinConfig.Secure
	rpcConfig.BufferSize = configuration.WorldpayWithinConfig.BufferSize

	rpcConfigReturn = rpcConfig

	return

}
