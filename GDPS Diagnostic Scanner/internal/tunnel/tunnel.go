package tunnel

import (
    "fmt"
    "net"
    "time"
)

func Test(addr string) string {
    conn, err := net.DialTimeout("udp", addr, 3*time.Second)
    if err != nil {
        return "Tunnel ERROR: " + err.Error()
    }
    conn.Close()
    return "Tunnel OK"
}
