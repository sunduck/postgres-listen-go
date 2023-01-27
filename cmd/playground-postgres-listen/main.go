package main

import (
	"context"
	"log"
	"time"

	"github.com/lib/pq"
)

// Notifying app thru postgress channels:
// q := fmt.Sprintf("SELECT pg_notify('%s', '%s')", 'eventChannel', 'eventName')
//	_, err := dls.db.Query(q)

func main() {
	channel := "eventChannel"
	conn := "host=db port=5432 user=postgres password=postgres dbname=playground_postgres_listen sslmode=disable"

	reportProblem := func(ev pq.ListenerEventType, err error) {
		if err != nil {
			log.Println(err)
		}
	}

	minReconn := 10 * time.Second
	maxReconn := time.Minute
	listener := pq.NewListener(conn, minReconn, maxReconn, reportProblem)
	defer listener.Close()

	if err := listener.Listen(channel); err != nil {
		log.Println(err)
		return
	}

	ctx := context.TODO()

	for {
		select {
		case <-ctx.Done():
			return
		case notice := <-listener.Notify:
			log.Printf("Notice with eventName: %s", notice.Extra)
		case <-time.After(90 * time.Second):
			go listener.Ping()
		}
	}
}
