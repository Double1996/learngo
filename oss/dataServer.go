package main

import (
	"github.com/double1996/learngo/oss/apiServer/heartbeat"
	"github.com/double1996/learngo/oss/apiServer/locate"
	"github.com/double1996/learngo/oss/apiServer/objects"
	"log"
	"net/http"
	"os"
)

func main() {

	go heartbeat.StartHeartbeat()
	go locate.StartLocate()
	http.HandleFunc("/objects/", objects.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
