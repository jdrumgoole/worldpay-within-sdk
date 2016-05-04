package client
import
(
	log "github.com/Sirupsen/logrus"
	"os"
)

func Main() {

	initLog()
}

func initLog() {

	log.SetFormatter(&log.JSONFormatter{})

	// Output to stderr instead of stdout, could also be a file.
	log.SetOutput(os.Stderr)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)

	log.Debug("Log is initialised. This is Worldpay Within Client app speaking.")
}
