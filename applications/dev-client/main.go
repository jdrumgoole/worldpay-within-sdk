package main
import
(

	"fmt"
	log "github.com/Sirupsen/logrus"
	"os"
	"time"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/utils/wslog"
)

func main() {

	initLog()

	doUI()

//	doWebSocketLogger()
}

func initLog() {

	log.SetFormatter(&log.JSONFormatter{})

	// Output to stderr instead of stdout, could also be a file.
	log.SetOutput(os.Stderr)

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