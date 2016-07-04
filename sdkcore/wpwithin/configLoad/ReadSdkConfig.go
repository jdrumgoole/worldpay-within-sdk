package configLoad

import (
    "fmt"
    "os"
    "encoding/json"
    rpc "innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/rpc"
)

type IndividualConfig struct {
    BufferSize  int
    Buffered    bool
    Framed      bool
    Host        string
    Logfile     string
    Loglevel    string
    Port        int
    Protocol    string
    Secure      bool
}

type ConfigurationLocal struct {
    WorldpayWithinConfig    IndividualConfig
}

func loadConfig() (configuration ConfigurationLocal) {
    file, _ := os.Open("conf.json")
    decoder := json.NewDecoder(file)
    configuration = ConfigurationLocal{}
    err := decoder.Decode(&configuration)
    if err != nil {
      fmt.Println("error:", err)
    }
    return
}

func PopulateConfiguration(rpcConfig rpc.Configuration) (rpcConfigReturn rpc.Configuration) {

    configuration := loadConfig()

    rpcConfig.Protocol = configuration.WorldpayWithinConfig.Protocol
    rpcConfig.Framed = configuration.WorldpayWithinConfig.Framed
    rpcConfig.Buffered = configuration.WorldpayWithinConfig.Buffered
    rpcConfig.Host = configuration.WorldpayWithinConfig.Host
    rpcConfig.Port = configuration.WorldpayWithinConfig.Port
    rpcConfig.Secure = configuration.WorldpayWithinConfig.Secure
    rpcConfig.BufferSize = configuration.WorldpayWithinConfig.BufferSize

    rpcConfigReturn = rpcConfig

    return

}

// func main() {

//     configuration := loadConfig()
//     fmt.Println(configuration.WorldpayWithinConfig)

// }