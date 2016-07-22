package main
import
(

	"fmt"
	log "github.com/Sirupsen/logrus"
	"time"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/utils/wslog"
	"os"
)

func main() {

	initLog()

	doUI()

//	doWebSocketLogger()
}

func initLog() {

	log.SetFormatter(&log.JSONFormatter{})

	f, err := os.OpenFile("output.log", os.O_WRONLY|os.O_CREATE, 0755)

	if err != nil {

		fmt.Println(err.Error())
	}

	log.SetOutput(f)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)

	log.Debug("Log is initialised.")
}

func doWebSocketLogger() {

	// Support all levels
	levels := make([]log.Level, 0)
	levels = append(levels, log.PanicLevel)
	levels = append(levels, log.FatalLevel)
	levels = append(levels, log.ErrorLevel)
	levels = append(levels, log.WarnLevel)
	levels = append(levels, log.InfoLevel)
	levels = append(levels, log.DebugLevel)

	err := wslog.Initialise("127.0.0.1", 8181, levels)

	if err != nil {

		fmt.Println(err.Error())

		return
	}
	
	time.Sleep(time.Duration(10) * time.Second)

	log.Debug("This is debug :)")
}