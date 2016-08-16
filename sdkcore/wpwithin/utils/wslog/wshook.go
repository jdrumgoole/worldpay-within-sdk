package wslog

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/websocket"
)

// WSHook A logrus hook to enable outputting logs to web browser via Web Socket
// Websocket implementation by Kevin Gordon Worldpay
type WSHook struct {
	ip       string
	port     int
	levels   []log.Level
	upgrader websocket.Upgrader
	wsConn   *websocket.Conn
}

// Initialise initialise the logger by specifing the IP and Port to listen on
// along with the levels to log out to
func Initialise(ip string, port int, levels []log.Level) error {

	hook := new(WSHook)
	hook.ip = ip
	hook.port = port
	hook.levels = levels
	hook.upgrader = websocket.Upgrader{} // Default options

	fmt.Println(hook.SocketClosedMsg())

	http.HandleFunc("/", hook.wsHome)
	http.HandleFunc("/connect", hook.wsConnect)
	listenAddr := fmt.Sprintf("%s:%d", hook.ip, hook.port)

	go func() {

		http.ListenAndServe(listenAddr, nil)
	}()

	log.AddHook(hook)

	return nil
}

// Levels avalable log levels
func (hook *WSHook) Levels() []log.Level {

	return hook.levels
}

// Fire fire an event
func (hook *WSHook) Fire(entry *log.Entry) error {

	var err error

	if hook.wsConn != nil {

		err = hook.wsConn.WriteMessage(websocket.TextMessage, []byte(entry.Message))
	} else {

		err = errors.New(hook.SocketClosedMsg())
	}

	return err
}

func (hook *WSHook) wsHome(w http.ResponseWriter, r *http.Request) {

	data := fmt.Sprintf("ws://%s/connect", r.Host)

	homeTemplate.Execute(w, data)
}

// SocketClosedMsg build a message notifying that the socket is closed
func (hook *WSHook) SocketClosedMsg() string {

	return fmt.Sprintf("Please open %s:%d in your browser and click Open to view logs", hook.ip, hook.port)
}

func (hook *WSHook) wsConnect(w http.ResponseWriter, r *http.Request) {

	conn, err := hook.upgrader.Upgrade(w, r, nil)

	if err != nil {

		fmt.Println(err.Error())
	}

	hook.wsConn = conn
	defer hook.wsConn.Close()

	for {

		time.Sleep(time.Duration(5 * time.Second))
	}
}
