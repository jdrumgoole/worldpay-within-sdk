package main

// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Edited by Kevin Gordon Worldpay

// +build ignore

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
	"time"
    "net"
)

var portNumber = "8181"

var addr = flag.String("addr", "localhost:" + portNumber, "http service address")

var upgrader = websocket.Upgrader{} // use default options

var c *websocket.Conn

func echo(w http.ResponseWriter, r *http.Request) {
	_c, err := upgrader.Upgrade(w, r, nil)

    // c gains local scope so need to use temporary variable to get global
    c = _c

	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {

		time.Sleep(time.Duration(5 * time.Second))

		if err != nil {
			log.Println("write:", err)
			break
		}

	}
}


func getIpAddress() string {
    var ipString string

    addrs, err := net.InterfaceAddrs()
    if err != nil {
        return "IP not detected"        
    }

    for _, a := range addrs {
        if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
            if ipnet.IP.To4() != nil {
                ipString = ipnet.IP.String()
                return ipString
            }
        }
    }

    return ipString
}

var socketClosedMsgShown = false

func showSocketClosedMsg() {
    var socketClosedMsg = "Please open " + getIpAddress() + ":" + portNumber + " in your browser and click Open to view logs";
    if(!socketClosedMsgShown) {
        log.Println(socketClosedMsg)
        socketClosedMsgShown = true;
    }
}
func EchoLogMsg(rs string) error {
    var err error  

    log.Println(rs)

    if c == nil {
        showSocketClosedMsg();
    } else {
        err = c.WriteMessage(websocket.TextMessage, []byte(rs))
        if err != nil {
           showSocketClosedMsg();
        }
    }

    return err
}



func home(w http.ResponseWriter, r *http.Request) {
	homeTemplate.Execute(w, "ws://"+r.Host+"/echo")
}

func entryPoint() {

    http.HandleFunc("/echo", echo)
    http.HandleFunc("/", home)
    log.Fatal(http.ListenAndServe(*addr, nil))

}

var homeTemplate = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<head>
<meta charset="utf-8">
<script>  
window.addEventListener("load", function(evt) {
    var output = document.getElementById("output");
    var input = document.getElementById("input");
    var ws;
    var print = function(message) {
        var d = document.createElement("div");
        d.innerHTML = message;
        output.appendChild(d);
    };
    document.getElementById("open").onclick = function(evt) {
        console.log("open");
        if (ws != null) {
            console.log("ws open so nothing further to do");
            return false;
        }
        ws = new WebSocket("{{.}}");
        ws.onopen = function(evt) {
            console.log("Should have openened ws");
            print("OPEN");
        }
        ws.onclose = function(evt) {
            console.log("Should now close ws");
            print("CLOSE");
            ws = null;
        }
        ws.onmessage = function(evt) {
            console.log("should message the ws");
            print("RESPONSE: " + evt.data);
        }
        ws.onerror = function(evt) {
            console.log("should ouput error on ws");
            print("ERROR: " + evt.data);
        }       
        console.log("should return from open onclick method");
        return false;
    };
    document.getElementById("justclose").onclick = function(evt) {
        console.log("justclose");
        if (!ws) {
             console.log("!ws so returning false");
             return false;
        }
        console.log("Actually closing ws");
        ws.close();
        ws.onclose();
        return false;
    };    
});
</script>
<style>
.mainFont {
    font-family: 'Century Gothic', CenturyGothic, AppleGothic, sans-serif;
}
.codeFont {
    font-family: 'Lucida Sans Typewriter', 'Lucida Console', monaco, 'Bitstream Vera Sans Mono', monospace;
    color: white;
    background-color: black;
}
</style>
</head>
<body class="mainFont">
<table>
<tr><td valign="top" width="50%">
<p class="mainFont">Click "Open" to create a connection to the Rapsberry Pi Log. Click "Close" to close the connection.
<p>
<button id="open">Open</button>
<button id="justclose">Close</button>
</td><td valign="top" width="50%">
<div id="output" class="codeFont"></div>
</td></tr></table>
</body>
</html>
`))