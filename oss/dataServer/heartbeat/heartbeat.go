package heartbeat

import (
	"github.com/double1996/learngo/oss/rabbitmq"
	"os"
	"time"
)

func StartHeartbeat()  {
	q := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
	defer q.Close()
	for {
		q.Publish("apiServers", os.Getenv("LISTEN_ADDRESS"))
		time.Sleep(5 * time.Second)
	}
}