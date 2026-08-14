package ports

import (
    "fmt"
    "net"
    "time"
)

func Check(port int, proto string) string {
    address := fmt.Sprintf("127.0.0.1:%d", port)

    conn, err := net.DialTimeout(proto, address, 3*time.Second)
    if err != nil {
        return fmt.Sprintf("Port %d (%s) ERROR: %s", port, proto, err.Error())
    }
    conn.Close()

    return fmt.Sprintf("Port %d (%s) OK", port, proto)
}
