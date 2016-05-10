package client
import
(
	"strconv"
	"os"
	"time"
	log "innovation.worldpay.com/worldpay-within-sdk/logrus"
)

func Main() {

	initLog()

	doStuff() // Replace with actually doing stuff!

}

func initLog() {


	log.SetFormatter(&log.JSONFormatter{})

	// Output to stderr instead of stdout, could also be a file.
	log.SetOutput(os.Stderr)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)

	log.Debug("Log is initialised.")

	log.SetSecureSocketsServer(true)

}




func doStuff() {

	var i = 0
	for {
		time.Sleep(time.Duration(2 * time.Second))
		log.Info("[" + strconv.Itoa(i) + "]")
		i++
	}

}