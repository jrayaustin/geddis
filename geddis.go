
package geddis

import "fmt"
import . "github.com/xuyu/goredis"
import "time"

var (
    network  = "tcp"
    address  = "127.0.0.1:6379"
    db       = 1
    password = ""
    timeout  = 5 * time.Second
    maxidle  = 1
    r        *Redis
    format = "tcp://auth:%s@%s/%d?timeout=%s&maxidle=%d"
)

type Geddis struct {}

type Geceiver struct {}

func (self Geddis) Listen() (*Geceiver, error) {

    client, err := Dial(&DialConfig{network, address, db, password, timeout, maxidle})

    if err != nil {
      return nil, err
    }

    quit := make(chan bool)
    sub, err := client.PubSub()
    defer sub.Close()

    go func() {
      if err := sub.Subscribe("go:event"); err != nil {
        quit <- true
      }
      for {
        list, err := sub.Receive()
        if err != nil {
          fmt.Println("an error")
          quit <- true
          break
        }
        // msgType := list[ 0 ]
        // topic := list[ 1 ]
        message := list[ 2 ]
        fmt.Println(message)
      }
    }()
    time.Sleep(100 * time.Millisecond)
    time.Sleep(100 * time.Millisecond)
    <-quit
    return &Geceiver{}, nil
}
