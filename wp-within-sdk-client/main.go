package main
import
(
	"fmt"
	"strconv"
	log "github.com/Sirupsen/logrus"
	"os"
	"time"
	"runtime"
	"sync"
)

func main() {

	//initLog()

	testWebSocketServer()
}

func initLog() {

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
		entryPoint()
	}()

	go func() {
		defer wg.Done()
		var i = 0
		for {
			time.Sleep(time.Duration(2 * time.Second))
			log.Debug("Should be outputting")
			EchoLogMsg("i = **" + strconv.Itoa(i) + "**")
			i++
		}
	}()	

	fmt.Println("Waiting To Finish")
    wg.Wait()
}