package ftp

import (
    "github.com/jlaffaye/ftp"
    "time"
)

func Check(addr string) string {
    conn, err := ftp.Dial(addr, ftp.DialWithTimeout(5*time.Second))
    if err != nil {
        return "FTP ERROR: " + err.Error()
    }
    conn.Quit()
    return "FTP OK"
}
