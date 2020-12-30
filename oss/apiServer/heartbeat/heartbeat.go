package heartbeat

import (
	"github.com/double1996/learngo/oss/rabbitmq"
	"os"
	"sync"
	"time"
)


var dataServers = make(map[string]time.Time)
var mutex sync.Mutex


func StartHeartbeat()  {
	q := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
	defer q.Close()
	for {
		q.Publish("apiServers", os.Getenv("LISTEN_ADDRESS"))
		time.Sleep(5 * time.Second)
	}
}

func ListenHeartbeat()  {
	q := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
	defer  q.Close()
	q.Bind("apiServer")
}