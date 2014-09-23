
package geddis

import "testing"
import "fmt"

func TestIt(t *testing.T) {
    g := Geddis{}
    geceiver, err := g.Listen()
    if err != nil {
		    t.Error(err)
	  }
    fmt.Println(geceiver)
}
