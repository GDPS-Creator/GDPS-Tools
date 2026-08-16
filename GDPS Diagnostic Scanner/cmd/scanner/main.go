package main

import (
    "fmt"
    "gdps-diagnostic-scanner/internal/mysql"
    "gdps-diagnostic-scanner/internal/ftp"
    "gdps-diagnostic-scanner/internal/ports"
    "gdps-diagnostic-scanner/internal/tunnel"
    "gdps-diagnostic-scanner/internal/report"
)

func main() {
    fmt.Println("GDPS Diagnostic Scanner started")

    results := report.New()

    results.Add(mysql.Check("localhost", "root", "password", "gdps"))
    results.Add(ftp.Check("127.0.0.1:21"))
    results.Add(ports.Check(4342, "udp"))
    results.Add(ports.Check(5351, "tcp"))
    results.Add(tunnel.Test("your.playit.gg.address:4342"))

    results.Print()
}
