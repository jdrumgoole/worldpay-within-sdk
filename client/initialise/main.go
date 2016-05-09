package client
import
(
	"fmt"
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
	var _wssKEV = &wsserver.Wsserver{};
	wssKEV = *_wssKEV
	testWebSocketServer()
	initLog()

}

func initLog() {

	log.SetSecureSocketsServer(wssKEV)

	log.SetFormatter(&log.JSONFormatter{})

	// Output to stderr instead of stdout, could also be a file.
	log.SetOutput(os.Stderr)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)

	log.Debug("Log is initialised.")
}

func testWebSocketServer() {
	log.Println("testWebSocketServer", nil)

	runtime.GOMAXPROCS(2)
	var wg sync.WaitGroup
    wg.Add(2)

	go func() {
		defer wg.Done()

		// debb kev - this creates a memory issue
		//if(wssKEV != nil) {
			log.Info("WebSocketServer EntryPoint calling");
			wssKEV.EntryPoint() 
			wssKEV.ShowSocketClosedMsg();
			log.Info("WebSocketServer EntryPoint called");

		//} else {
		// 	log.Info("WebSocketServer was nil so couldn't start");
		//}

	}()

	go func() {
		defer wg.Done()
		var i = 0
		for {
			time.Sleep(time.Duration(2 * time.Second))
			log.Debug("Should be outputting")
			log.Info("This is via log.info: i = **" + strconv.Itoa(i) + "** ANOTHER CHANGE")
			i++
		}
	}()	

	fmt.Println("Waiting To Finish")
    wg.Wait()


}