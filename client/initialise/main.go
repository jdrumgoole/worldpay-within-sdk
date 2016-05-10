package client
import
(
	"strconv"
	"os"
	"time"
	"runtime"
	"sync"
	log "github.com/Sirupsen/logrus"
	wsserver "innovation.worldpay.com/worldpay-within-sdk/client/Wsserver"
)

var wssKEV wsserver.Wsserver

func Main() {

	initLog()

}

func initLog() {


	log.SetFormatter(&log.JSONFormatter{})

	// Output to stderr instead of stdout, could also be a file.
	log.SetOutput(os.Stderr)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)

	log.Debug("Log is initialised.")

	setupServer();

	log.SetSecureSocketsServer(wssKEV)

}

func setupServer() {
	var _wssKEV = &wsserver.Wsserver{};
	wssKEV = *_wssKEV
	startWebSocketServer()
}

func startWebSocketServer() {

	runtime.GOMAXPROCS(2)
	var wg sync.WaitGroup
    wg.Add(2)

	go func() {

		defer wg.Done()
		log.Info(wssKEV.ShowSocketClosedMsg());		
		wssKEV.EntryPoint() 

	}()

	go func() {
		defer wg.Done()
		var i = 0
		for {
			time.Sleep(time.Duration(2 * time.Second))
			log.Info("[" + strconv.Itoa(i) + "]")
			i++
		}
	}()	

	wg.Wait()

}